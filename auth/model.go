package auth

import "github.com/oklog/ulid/v2"

// Auth ...
type Auth struct {
	ID ulid.ULID
	// AccountID ulid.ULID
	Token string
}
