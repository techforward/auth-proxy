package account

import "github.com/oklog/ulid/v2"

type Account struct {
	ID           ulid.ULID
	Name         string
	PasswordHash string
}
