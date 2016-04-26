package password_test

import (
	"bytes"

	. "github.com/hyperboloide/users/password"
	"golang.org/x/crypto/bcrypt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testPassword struct {
	hash    []byte
	version int
}

func (tp *testPassword) PasswordBytes() []byte {
	return tp.hash
}

func (tp *testPassword) PasswordVersion() int {
	return tp.version
}

func (tp *testPassword) PasswordSet(h []byte, v int) error {
	tp.hash = h
	tp.version = v
	return nil
}

var _ = Describe("Password", func() {

	passwd := &testPassword{}

	It("should update and check a password", func() {
		Ω(Update(passwd, "secret")).To(BeNil())
		Ω(passwd.version).To(Equal(1))
		Ω(passwd.hash).ToNot(BeNil())

		ok, err := Check(passwd, "wrong")
		Ω(ok).To(BeFalse())
		Ω(err).To(BeNil())

		ok, err = Check(passwd, "secret")
		Ω(ok).To(BeTrue())
		Ω(err).To(BeNil())
	})

	It("should increment version", func() {
		CurrentVersion = 2
		CurrentHasher = Bcrypt{bcrypt.MinCost}
		Previous = map[int]Hasher{
			1: Bcrypt{14},
		}

		save := passwd.hash
		ok, err := Check(passwd, "secret")
		Ω(ok).To(BeTrue())
		Ω(err).To(BeNil())
		Ω(passwd.version).To(Equal(2))
		Ω(bytes.Equal(passwd.hash, save)).To(BeFalse())
	})

})
