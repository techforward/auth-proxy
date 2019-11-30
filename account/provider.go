package account

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// Provider ...
type Provider struct {
	r Repository
}

// NewProvider ...
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

func (p *Provider) create(username, password string) Account {
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

func (p *Provider) update() {}

func (p *Provider) delete() {}

func (p *Provider) changePassword() {}
