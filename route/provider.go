package route

// Provider property
type Provider struct {
	mapping map[string]string
}

var domainHost = "lvh.me"

// New ...
func New() Repository {
	mapping := make(map[string]string)
	// mapping["jwt-auth."+domainHost] = "172.0.0.1:8080"
	// mapping[domainHost] = domainHost
	mapping["cat-fact."+domainHost] = "cat-fact.herokuapp.com"
	mapping["api-dot-soccer-250403."+domainHost] = "api-dot-soccer-250403.appspot.com"
	return &Provider{mapping}
}

// GetDest ...
func (p *Provider) GetDest(src string) string {
	if dest, ok := p.mapping[src]; ok {
		return dest
	}
	return src
}

// AddRoute ...
func (p *Provider) AddRoute() bool { return true }
