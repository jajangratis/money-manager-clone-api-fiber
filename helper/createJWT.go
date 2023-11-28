package helper

import (
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(key string, value string) (string, error) {
	godotenv.Load()
	// Create a new token object, specifying signing method and claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set claims (payload)
	claims[key] = value
	//claims["name"] = "John Doe"
	//claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix() // Token expiration time

	// Secret key for signing the token

	secretKey, _ := base64.StdEncoding.DecodeString(os.Getenv("KEY_JWT"))
	//secretKey := []byte(os.Getenv("KEY_JWT"))
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ExtractJWTToken(tokenString string) (jwt.MapClaims, error) {
	godotenv.Load()
	// Decode the base64-encoded string
	decodedBytes, err := base64.StdEncoding.DecodeString(os.Getenv("KEY_JWT"))
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// You should validate and return the key used to sign the token
		// For simplicity, we're using a sample key here. In production, use a secure method to retrieve the key.
		//return []byte(os.Getenv("KEY_JWT")), nil
		return []byte(decodedBytes), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid JWT token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims from JWT token")
	}

	return claims, nil
}

//func main() {
//	// Create and sign a JWT
//	jwtString, err := createJWT()
//	if err != nil {
//		fmt.Println("Error creating JWT:", err)
//		return
//	}
//
//	fmt.Println("JWT:", jwtString)
//}
