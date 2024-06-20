package main

func main() {
}

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	usersMap := make(map[string]struct{}, len(users))
	uniqueUsers := []User{}

	for _, i := range users {
		_, ok := usersMap[i.Nickname]
		if !ok {
			uniqueUsers = append(uniqueUsers, i)
			usersMap[i.Nickname] = struct{}{}
		}
	}

	return uniqueUsers
}
