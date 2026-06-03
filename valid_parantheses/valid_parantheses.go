package main

import "fmt"

var parantheses_map = map[string]string{
	"]": "[",
	"}": "{",
	")": "(",
}

func check_valid_parantheses(sequence string) bool {
	var hash_array []string
	for _, char := range sequence {
		expected_paranthesis, exists := parantheses_map[string(char)]
		if len(hash_array) == 0 && exists {
			return false
		} else if exists && hash_array[len(hash_array)-1] == expected_paranthesis {
			hash_array = hash_array[:len(hash_array)-1]
		} else {
			hash_array = append(hash_array, string(char))
		}
	}
	var is_empty bool = len(hash_array) == 0
	return is_empty
}

func main() {
	sequences := [...]string{"()[]", "{}()[]", ")[]", "()[", "[](){}[])"}
	for _, sequence := range sequences {
		is_valid := check_valid_parantheses(sequence)
		if is_valid {
			fmt.Printf("%s is valid\n", sequence)
		} else {
			fmt.Printf("%s is not valid\n", sequence)
		}

	}
}
