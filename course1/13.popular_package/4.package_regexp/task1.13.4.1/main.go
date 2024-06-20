package main

import (
	"regexp"
)

func main() {
}

func isValidEmail(email string) bool {
	mail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&*+/=?^_{|}~-]+@[a-zA-Z0-9-]+[a-zA-Z/.]{2,}$")
	return mail.MatchString(email)
}
