package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "Archawat",
		Email:      "M_zoman@hotmail.com",
		CustomerID: "M1234567",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("Customerid not correct format"))
}
