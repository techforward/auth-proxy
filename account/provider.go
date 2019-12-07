package account

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type provider struct{}

// New ...
func New() Repository {
	return &provider{}
}

func (p *provider) create(username, password string) Account {
	t := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ulid := ulid.MustNew(ulid.Timestamp(t), entropy)

	account := &Account{
		ulid,
		username,
		password,
	}
	return *account
}

func (p *provider) update() Account { return Account{} }

func (p *provider) delete() bool { return true }

func (p *provider) changePassword() bool { return true }
