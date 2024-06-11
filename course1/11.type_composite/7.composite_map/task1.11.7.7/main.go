package main

func main() {
}

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	usersMap := make(map[string]bool)
	uniqueUsers := []User{}
	for _, i := range users {
		if !usersMap[i.Nickname] {
			uniqueUsers = append(uniqueUsers, i)
			usersMap[i.Nickname] = true
		}
	}
	answer := make([]User, len(uniqueUsers))
	copy(answer, uniqueUsers)
	return answer
}
