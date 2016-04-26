package mock

import "github.com/dchest/uniuri"

// User is a user mockup.
type User struct {
	id       string
	login    string
	email    string
	language string
}

// GetID returns an id.
func (u *User) GetID() string { return u.id }

// GetLogin returns a login.
func (u *User) GetLogin() string { return u.login }

// GetEmail returns an email.
func (u *User) GetEmail() string { return u.email }

// GetLanguage returns a language.
func (u *User) GetLanguage() string { return u.language }

// Provider provides mock Users objects.
type Provider struct{}

// Generate a new mock user
func Generate() *User {
	return &User{
		uniuri.NewLen(6),
		uniuri.NewLen(6),
		uniuri.NewLen(6) + "@" + uniuri.NewLen(10) + ".com",
		"en",
	}
}

// ByID generates an new mock User.
func (p Provider) ByID(string) (*User, error) { return Generate(), nil }

// ByLogin generates an new mock User.
func (p Provider) ByLogin(string) (*User, error) { return Generate(), nil }

// ByEmail generates an new mock User.
func (p Provider) ByEmail(string) (*User, error) { return Generate(), nil }
