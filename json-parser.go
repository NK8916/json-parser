package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	is_valid := isValidJson(args)
	if is_valid {

	} else {
		os.Exit(1)
	}
}

func isValidJson(json_string_arr []string) bool {
	if len(json_string_arr) != 1 {
		return false
	}
	json_string := json_string_arr[0]
	stack := []string{}
	for _, char := range json_string {
		if char == '{' {
			stack = append(stack, string(char))
		} else if char == '}' && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
