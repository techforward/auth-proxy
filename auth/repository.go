package auth

// Repository auth
type Repository interface {
	CreateToken() string
	VerifyToken(string) bool
}
