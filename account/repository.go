package account

// Repository ...
type Repository interface {
	create(string, string) Account
	update() Account
	delete() bool
	changePassword() bool
}
