// exercise 3.10
package main

import (
	"bytes"
	"fmt"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		return s
	}
	remainder := len(s) % 3
	buf.WriteString(s[:remainder])
	s = s[remainder:]
	for len(s) > 0 {
		if remainder > 0 {
			buf.WriteString(",")
		} else {
			remainder = 1
		}

		buf.WriteString(s[:3])
		s = s[3:]
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("12345"))
	fmt.Println(comma("1000000000"))
	fmt.Println(comma("123451234512345"))
	fmt.Println(comma("12345123"))
}
		
