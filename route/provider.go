package route

// Provider ...
type Provider struct {
	r Repository
}

// NewProvider ...
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

func (p *Provider) getDest(src string) string {
	return "localhost:3000"
}