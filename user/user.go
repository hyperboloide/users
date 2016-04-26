package user

// User allows access to a user data.
type User interface {
	GetID() string
	GetLogin() string
	GetEmail() string
	GetLanguage() string
}

// Provider provides User objects.
type Provider interface {
	ByID(string) (*User, error)
	ByLogin(string) (*User, error)
	ByEmail(string) (*User, error)
}
