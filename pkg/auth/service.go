package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Authentication errors
var (
	ErrUnauthorised = errors.New("user is unauthorised")
)

// Authenticator handles authentication and verification on accessing user
type Authenticator interface {
	// Grant valid user an access token
	// Returns corresponding error if rejected
	Grant(user, password string) (string, error)

	// Verify a token is still valid or not
	Verify(tokenStr string) (bool, error)
}

// Service is an authentication service for RESTful API
type Service struct {
	store     CredentialStore
	jwtSecret []byte
}

// NewDummyService is a dummy authentication service whichs purposes is for coding example
func NewDummyService() Authenticator {
	return &Service{
		store:     NewMockCredentialStore(),
		jwtSecret: []byte("hardcoded_private_key"), // Hardcoded key only lives in coding example
	}
}

// Grant valid user an access token
// Returns a new access token
// Returns corresponding error if rejected
func (s *Service) Grant(user, password string) (string, error) {
	if s.store.Check(user, password) {
		jwtLifetime := time.Now().Add(30 * time.Minute) // The JWT will only live 30 minutes
		claims := jwt.MapClaims{
			"username": user,
			"exp":      jwtLifetime.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString(s.jwtSecret)

		return tokenStr, err
	}

	return "", ErrUnauthorised
}

// Verify a token is still valid or not
// Returns corresponding error if token is rejected
func (s *Service) Verify(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.jwtSecret, nil
	})
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, ErrUnauthorised // Token is not valid
}
