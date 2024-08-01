package entities

type Book struct {
	Id      int
	Name    string
	Author  *Author
	TakenBy *User
}

type User struct {
	Id         int
	UserName   string
	BooksTaken []*Book
}

type Author struct {
	Id    int
	Name  string
	Books []*Book
}
