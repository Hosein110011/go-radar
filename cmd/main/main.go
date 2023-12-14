package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/Hosein110011/go-radar/pkg/routes"
)


func main() {
	erre := godotenv.Load(".env") // Load .env file
	if erre != nil {
		log.Fatal("Error loading .env file", erre)
	}

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	portStr := os.Getenv("PORT")
	port, errr := strconv.Atoi(portStr)
	if errr != nil {
		log.Fatalf("Invalid port number: %v\n", errr)
	}

	serverAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("WebSocket Server listening on http://localhost%s\n", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Println("Error:", err)
	}
}