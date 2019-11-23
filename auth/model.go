package auth

import "github.com/oklog/ulid/v2"

type Auth struct {
	ID ulid.ULID
	// AccountID ulid.ULID
	Token string
}
