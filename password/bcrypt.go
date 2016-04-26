package password

import "golang.org/x/crypto/bcrypt"

// Bcrypt implements a bcrypt Hasher
type Bcrypt struct {
	//If the cost given is 0 or less than bcrypt.MinCost then bcrypt.Default
	// will be used.
	Cost int
}

// Hash a []byte hash from a string
func (b Bcrypt) Hash(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
}

// Compare a []byte hash to a string
func (b Bcrypt) Compare(h []byte, p string) bool {
	return bcrypt.CompareHashAndPassword(h, []byte(p)) == nil
}
