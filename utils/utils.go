package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"semay.com/config"
)

var key = config.Config("TOKEN_SALT")

type UserClaim struct {
	jwt.RegisteredClaims
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	UUID  string   `json:"uuid"`
}

// Generate 16 bytes randomly
func GenerateSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

// Combine password and salt then hash them using the SHA-512
func HashFunc(password string) string {

	// var salt []byte
	// get salt from env variable
	salt := []byte(config.Config("SECRETE_SALT"))

	// Convert password string to byte slice
	var pwdByte = []byte(password)

	// Create sha-512 hasher
	var sha512 = sha512.New()

	pwdByte = append(pwdByte, salt...)

	sha512.Write(pwdByte)

	// Get the SHA-512 hashed password
	var hashedPassword = sha512.Sum(nil)

	// Convert the hashed to hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPassword)
	return hashedPasswordHex
}

func PasswordsMatch(hashedPassword, currPassword string) bool {

	var currPasswordHash = HashFunc(currPassword)

	return hashedPassword == currPasswordHash
}

// source of this token encode decode functions
// https://github.com/gurleensethi/go-jwt-tutorial/blob/main/main.go
func CreateJWTToken(email string, uuid string, roles []string, duration int) (string, error) {
	my_claim := UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{},
		Email:            email,
		Roles:            roles,
		UUID:             uuid,
	}

	exp := time.Now().Add(time.Duration(duration) * time.Minute)
	my_claim.ExpiresAt = jwt.NewNumericDate(exp)
	my_claim.Issuer = "Blue Admin"
	my_claim.Subject = "UI Authentication Token"
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, my_claim)

	signedString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}

	return signedString, nil
}

func ParseJWTToken(jwtToken string) (map[string]interface{}, error) {

	response_claim := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtToken, response_claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	// check token validity, for example token might have been expired
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return response_claim, nil

}
