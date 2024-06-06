package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func main() {
	fmt.Println(GenerateFakeData())
}

func GenerateFakeData() string {
	fakeName := fake.FullName()
	fakeAdress := fake.StreetAddress()
	fakePhone := fake.Phone()
	fakeMail := fake.EmailAddress()
	return fmt.Sprint("Name: ", fakeName, "\nAddress: ", fakeAdress, "\nPhone: ", fakePhone, "\nEmail: ", fakeMail)
}
