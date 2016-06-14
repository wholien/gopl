package main

import (
	"fmt"
)

func anagram(a, b string) bool {
	if len(a) != len(b) {return false}

	var counter [26]int
	for i := 0; i < len(a); i++ {
		counter[int(a[i]-'a')]++
	}
	for i := 0; i < len(b); i++ {
		if counter[int(b[i]-'a')] == 0 {
			return false
		}
		counter[int(b[i]-'a')]--
	}
	return true
}

func main() {
	fmt.Println(anagram("asdf", "asfd"))
	fmt.Println(anagram("asdf", "asdfd"))
	fmt.Println(anagram("asdf", "asdg"))
	fmt.Println(anagram("asdfasdfasdf", "fdsafdsafdsa"))
}
