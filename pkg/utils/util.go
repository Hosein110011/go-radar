package utils

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
)


type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// Function to extract and parse the token
func GetTokenFromHeader(r *http.Request) (*jwt.Token, error) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        return nil, fmt.Errorf("authorization header is not provided")
    }

    // Split the header to get the token part
    headerParts := strings.Split(authHeader, " ")
    if len(headerParts) != 2 || headerParts[0] != "Bearer" {
        return nil, fmt.Errorf("authorization header format must be 'Bearer {token}'")
    }

    // Parse the token
    token, err := jwt.ParseWithClaims(headerParts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
        // Here, you need to return the key used to sign the token
        // For example, return []byte("your-256-bit-secret"), nil
        return []byte("django-insecure-l2l5p!b)h!^pe0^im!yp!hkv$anb*jkx$o-_$y3vc%e8%=$(@9"), nil
    })

    return token, err
}

