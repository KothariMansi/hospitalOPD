package token

import "time"

// Maker is an interface for managing token
type Maker interface {
	// CreateToken create a new token for specific username ans duration
	CreateToken(username string, duration time.Duration) (string, error)
	// VerifyToken checks if the token id is valid or not
	VerifyToken(token string) (*Payload, error)
}
