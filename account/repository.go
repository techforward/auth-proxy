package account

type Repository interface {
	create() Account
	update() Account
	delete() bool
	changePassword() bool
}
