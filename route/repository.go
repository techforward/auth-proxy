package route

type Repository interface {
	verifyHost(domainHosts map[string]bool, host string) bool
	addRoute()
	getDest(scr string) string
}
