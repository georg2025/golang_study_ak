package main

type User struct {
	ID   int
	Name string
	Age  int
}

func Merge(arr1 []User, arr2 []User) []User {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	result := make([]User, 0, (len(arr1) + len(arr2)))
	arr1Flag, arr2Flag := 0, 0

	for arr1Flag < len(arr1) && arr2Flag < len(arr2) {
		if arr1[arr1Flag].ID < arr2[arr2Flag].ID {
			result = append(result, arr1[arr1Flag])
			arr1Flag += 1
		} else {
			result = append(result, arr2[arr2Flag])
			arr2Flag += 1
		}
	}

	for arr1Flag < len(arr1) {
		result = append(result, arr1[arr1Flag])
		arr1Flag += 1
	}

	for arr2Flag < len(arr2) {
		result = append(result, arr2[arr2Flag])
		arr2Flag += 1
	}

	return result
}
