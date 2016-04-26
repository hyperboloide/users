package user

// User allows access to a user data.
type User interface {
	GetLogin() string
	GetEmail() string
}

// Provider provides User objects.
type Provider interface {
	ByLogin(string) (User, error)
	ByEmail(string) (User, error)
}
