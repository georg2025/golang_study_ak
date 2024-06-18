package main

func main() {
}

func getType(i interface{}) string {
	if i == nil {
		return "Empty interface"
	}
	answer := ""
	switch i.(type) {
	case int:
		answer = "int"
	case string:
		answer = "string"
	case []int:
		answer = "[]int"
	default:
		answer = "cant figure out"
	}
	return answer
}
