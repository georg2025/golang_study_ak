package main

import (
	"strings"
)

func main() {
}

type User struct {
	Name     string
	Comments []Comment
}

type Comment struct {
	Message string
}

func FilterComments(users []User) []User {
	newUsers := []User{}
	for _, i := range users {
		newComments := []Comment{}
		for _, j := range i.Comments {
			if !IsBadComment(j.Message) {
				newComments = append(newComments, j)
			}
		}
		i.Comments = newComments
		newUsers = append(newUsers, i)
	}
	return newUsers
}

func IsBadComment(comment string) bool {
	return strings.Contains(strings.ToLower(comment), "bad comment")
}

func GetBadComments(users []User) []Comment {
	answer := []Comment{}
	for _, i := range users {
		for _, j := range i.Comments {
			if strings.Contains(strings.ToLower(j.Message), "bad comment") {
				answer = append(answer, j)
			}
		}
	}
	return answer
}

func GetGoodComments(users []User) []Comment {
	answer := []Comment{}
	for _, i := range users {
		for _, j := range i.Comments {
			if strings.Contains(strings.ToLower(j.Message), "good comment") {
				answer = append(answer, j)
			}
		}
	}
	return answer
}
