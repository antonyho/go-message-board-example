package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestService_Grant(t *testing.T) {
	type testCase struct {
		Name          string
		Username      string
		Password      string
		ExpectToken   bool
		ExpectedError error
	}

	testCases := []testCase{
		{
			Name:          "ValidCredential",
			Username:      "admin",
			Password:      "back-challenge",
			ExpectToken:   true,
			ExpectedError: nil,
		},
		{
			Name:          "WrongCredential",
			Username:      "admin",
			Password:      "wrong-password",
			ExpectToken:   false,
			ExpectedError: ErrUnauthorised,
		},
	}

	authService := NewDummyService(30 * time.Minute)

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			tokenStr, err := authService.Grant(c.Username, c.Password)
			if c.ExpectedError != nil {
				assert.EqualError(t, err, c.ExpectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			if c.ExpectToken {
				assert.NotEmpty(t, tokenStr)
			} else {
				assert.Empty(t, tokenStr)
			}
		})
	}
}

func TestService_Verify(t *testing.T) {
	type testCase struct {
		Name              string
		TokenLifeDuration time.Duration
		Username          string
		Password          string
		CustomToken       string
		ExpectedResult    bool
		ExpectedError     error
	}

	testCases := []testCase{
		{
			Name:              "ValidToken",
			TokenLifeDuration: 30 * time.Minute,
			Username:          "admin",
			Password:          "back-challenge",
			ExpectedResult:    true,
			ExpectedError:     nil,
		},
		{
			Name:              "ExpiredToken",
			TokenLifeDuration: 1 * time.Nanosecond,
			Username:          "admin",
			Password:          "back-challenge",
			ExpectedResult:    false,
			ExpectedError:     errors.New("Token is expired"),
		},
		{
			Name:              "AlteredToken",
			TokenLifeDuration: 30 * time.Minute,
			Username:          "admin",
			Password:          "back-challenge",
			CustomToken:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU",
			ExpectedResult:    false,
			ExpectedError:     ErrUnauthorised,
		},
	}

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			authService := NewDummyService(c.TokenLifeDuration)
			tokenStr, err := authService.Grant(c.Username, c.Password)
			assert.NoError(t, err)

			time.Sleep(1 * time.Second)

			if c.CustomToken != "" {
				tokenStr = c.CustomToken
			}
			valid, err := authService.Verify(tokenStr)
			if c.ExpectedError != nil {
				assert.EqualError(t, err, c.ExpectedError.Error())
			}
			assert.Exactly(t, c.ExpectedResult, valid)
		})
	}
}
