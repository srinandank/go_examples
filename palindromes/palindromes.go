package main

import "fmt"

func check_palindrome(word string) bool {
	first, last := 0, len(word)-1
	for first < last {
		if word[first] != word[last] {
			return false
		}
		first++
		last--
	}
	return true
}

func main() {
	words := [...]string{"malayalam", "level", "jack"}

	for _, word := range words {
		var is_palindrome bool = check_palindrome(word)
		if is_palindrome {
			fmt.Printf("%s is a palindrome\n", word)
		} else {
			fmt.Printf("%s is not a palindrome\n", word)
		}
	}
}
