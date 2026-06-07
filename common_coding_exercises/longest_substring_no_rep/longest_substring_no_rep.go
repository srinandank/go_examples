package main

import "fmt"

func longest_substring(sequence string) (string, int) {
	first, last := 0, 1
	char_map := make(map[string]int)
	char_map[string(sequence[first])] = first
	longest_map := make(map[string]int)
	var length int = 1
	for last != len(sequence) {
		_, exists := char_map[string(sequence[last])]
		if exists {
			longest_map[string(sequence[first:last])] = length
			char_map[string(sequence[last])] = last
			length = 1
			first = last
		} else {
			length++
		}
		last++
	}
	var longest_sequence string
	var longest_len int
	for seq, seq_len := range longest_map {
		if longest_len < seq_len {
			longest_sequence = seq
			longest_len = seq_len
		}
	}
	return longest_sequence, longest_len
}

func main() {
	sequences := [...]string{"abcabc", "abcdea", "bbbbb", "zxyapz", "aaabcdeaaa"}

	for _, sequence := range sequences {
		longest_sequence, length_seq := longest_substring(sequence)
		fmt.Printf("Longest sequence of %s is %s of length %d\n", sequence, longest_sequence, length_seq)
	}
}
