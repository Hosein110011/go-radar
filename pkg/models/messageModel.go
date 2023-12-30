package models

import (
	"log"
	"sort"
	"time"

	"github.com/gocql/gocql"
)

type Message struct {
	OwnerId     string
	RecipientId string
	IsDeleted   bool
	Created     time.Time
	// Add other fields as needed
}

func GetSession() (session *gocql.Session) {
	cluster := gocql.NewCluster("cassandra") // replace with your Cassandra node IPs
	cluster.Keyspace = "cassandra"           // replace with your keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}
	return session
}

func GetSortedMessages(session *gocql.Session, ownerID, recipientID string, sliced int) []Message {
	var messages []Message

	// Query messages from owner to recipient
	query := "SELECT * FROM message WHERE owner_id = ? AND recipient_id = ? AND is_deleted = false LIMIT ? ALLOW FILTERING"
	iter := session.Query(query, ownerID, recipientID, sliced).Iter()
	var msg Message
	for iter.Scan(&msg.OwnerId, &msg.RecipientId, &msg.IsDeleted, &msg.Created /*, other fields */) {
		messages = append(messages, msg)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	// Query messages from recipient to owner (repeat the process as above)

	// Sort messages by creation time
	sort.SliceStable(messages, func(i, j int) bool {
		return messages[i].Created.After(messages[j].Created)
	})
	defer session.Close()

	return messages
}
