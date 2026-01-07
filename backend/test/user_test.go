package test

import(
	"testing"
	"github.com/yanat053/final-test/backend/entity"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func TestUserValidation (t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Test valid user", func(t *testing.T) {
		user := entity.User{
			Username: "validuser",
			Email:    "yanatchara@gmail.com",
			Password: "validpassword",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}