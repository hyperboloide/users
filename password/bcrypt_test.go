package password_test

import (
	. "github.com/hyperboloide/users/password"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bcrypt", func() {

	It("should hash a string", func() {
		hasher := Bcrypt{10}
		h, err := hasher.Hash("secret password")
		Ω(err).To(BeNil())
		Ω(h).ToNot(BeNil())

		Ω(hasher.Compare(h, "secret password")).To(BeTrue())
	})

})
