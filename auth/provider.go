package auth

import (
	"fmt"
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

func (p *Provider) createToken() {
	t := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	fmt.Println(ulid.MustNew(ulid.Timestamp(t), entropy))
}
