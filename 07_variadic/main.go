package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	varidicStrings(s...) // ส่งแบบ variadic
}

/// ...string คือการบอกว่าจะรับแบบ variadic
func varidicStrings(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println("index: ", i, " is : ", s[i])
	}
}
