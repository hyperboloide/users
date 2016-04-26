package password

import "fmt"

// Hasher is an interface implements a method to generate a Hash from a string
// and a method to compare a Hash to a string
type Hasher interface {
	Hash(string) ([]byte, error)
	Compare([]byte, string) bool
}

var (
	// CurrentVersion of the Hasher used, if the versions do not match
	// the previous function is used and replaced by a
	// hash from the current Hasher.
	CurrentVersion int

	// CurrentHasher is the default Hasher to be used.
	CurrentHasher Hasher

	// Previous are used when the Password do not match the
	// CurrentVersion.
	Previous map[int]Hasher
)

func init() {
	CurrentVersion = 1
	CurrentHasher = Bcrypt{14}
	Previous = map[int]Hasher{}
}

// Password allows acces to the user's password.
type Password interface {
	Bytes() []byte
	Version() int
	Set([]byte, int) error
}

// Check a Password againts a string.
func Check(password Password, provided string) (bool, error) {
	version := password.Version()
	hasher := CurrentHasher
	if version != CurrentVersion {
		if h, ok := Previous[version]; ok {
			hasher = h
		} else {
			return false, fmt.Errorf("Unknow hasher version '%d'", version)
		}
	}
	if !hasher.Compare(password.Bytes(), provided) {
		return false, nil
	} else if version == CurrentVersion {
		return true, nil
	}
	return true, Update(password, provided)
}

// Update the password with the provided string
func Update(password Password, provided string) error {
	h, err := CurrentHasher.Hash(provided)
	if err != nil {
		return err
	}
	return password.Set(h, CurrentVersion)
}
