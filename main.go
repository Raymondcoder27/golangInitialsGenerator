package main

import (
	"fmt"
	"strings"
)

// function to take in a name and return initials

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
func main() {
	//test the function
	fn, sn := getInitials("Mwebe Raymond")
	fn2, sn2 := getInitials("Ronald")
	fmt.Println(fn, sn)
	fmt.Println(fn2, sn2)
}
