package route

// Repository ...
type Repository interface {
	AddRoute() bool
	GetDest(string) string
}
