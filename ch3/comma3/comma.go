package main

import (
	"fmt"
	"strings"
)

func wrapper(s string) string {
	copy := s
	if copy[0] == '+' || copy[0] == '-' {
		copy = copy[1:]
		return string(s[0]) + wrapper(copy)
	}
	if i := strings.Index(copy, ".") ; i != -1 {
		copy = copy[:i]
		return wrapper(copy) + s[i:] 
	}
	return comma(s)
}
	
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(wrapper("12345"))
	fmt.Println(wrapper("+12345"))
	fmt.Println(wrapper("-12345"))
	fmt.Println(wrapper("+12345.123"))
	fmt.Println(wrapper("+12345.1231231234"))
}
