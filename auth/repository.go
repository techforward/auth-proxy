package auth

type Repository interface {
	createToken()
	verifyToken()
}
