package main

import "strings"

func Spacify(s string) string {
	s = strings.Replace(s, " ", "", -1)
	str := strings.Split(s, "")
	return strings.Join(str, " ")
}

func main() {
	Spacify("Pippi")
}
