package token

import (
	"github.com/google/uuid"
	"time"
)

// Maker is an interface for mapping token
type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(userID, cartID uuid.UUID, duration time.Duration) (string, *Payload, error)

	// VerifyToejn checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
