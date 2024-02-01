package unit

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/poon16015/prac2/entity"
)

func TestAll(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`test All`, func(t *testing.T){
		user := entity.User{
			Name : "",
			Email : "poon@gmail.com",
			Phone : "0123456789" ,
		}
		ok , err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
	})

}

func TestUserPhone(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`phone_number is required`, func(t *testing.T) {
		user := entity.User{
			Name: "poon", 
			Email: "poob.aa@gmail.com" ,// ผิดตรงนี้
			Phone :"",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("phone is required"))
	})
	
	t.Run(`phone_number check 10 digit`, func(t *testing.T) {
		user := entity.User{
			Name: "poon", 
			Email: "poob.aa@gmail.com" ,
			Phone :"01234567895555",// ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("PhoneNumber length is not 10 digits.")))

	})
}

func TestUserEmail(t *testing.T) {
	g := NewGomegaWithT(t)


		t.Run(`Url is required`, func(t *testing.T) {
		user := entity.User{
			Name: "poon", 
			Email: "" ,// ผิดตรงนี้
			Phone :"01234567895555",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Email is required")))
	})

	t.Run(`“PhoneNumber length is not 10 digits.”`, func(t *testing.T) {
		user := entity.User{
			Name: "poon", 
			Email: "poob.aa.gmail.com" ,// ผิดตรงนี้
			Phone :"01234567895555",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Email is invalid")))
	})
}