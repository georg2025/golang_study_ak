package main

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}
type UserOption func(*User)

func WithUsername(name string) UserOption {
	return func(u *User) {
		u.Username = name
	}
}

func WithEmail(email string) UserOption {
	return func(u *User) {
		u.Email = email
	}
}

func WithRole(role string) UserOption {
	return func(u *User) {
		u.Role = role
	}
}

func NewUser(id int, options ...UserOption) *User {
	user := &User{ID: id}

	for _, option := range options {
		option(user)
	}

	return user
}

func main() {
}
