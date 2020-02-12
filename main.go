package main

import (
	"fmt"

	"github.com/teohrt/kPalindromes/kpalindromes"
)

// Given an int k, and a string s, write a function that returns a bool
// representing if k number of palindromes can be created from
// permutations of s that utilize ALL characters of s.

func main() {
	k := 3
	s := "aabbaa"
	res := kpalindromes.Kpalindromes(k, s)
	fmt.Println(res)
}
