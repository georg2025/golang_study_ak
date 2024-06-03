package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func main() {
	fmt.Println(GenerateFakeData())
}

func GenerateFakeData() string {
	fakename := fake.FullName()
	fakeadress := fake.StreetAddress()
	fakephone := fake.Phone()
	fakemail := fake.EmailAddress()
	return fmt.Sprint("Name: ", fakename, "\nAddress: ", fakeadress, "\nPhone: ", fakephone, "\nEmail: ", fakemail)
}
