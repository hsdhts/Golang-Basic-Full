package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//func blacklistAdmin(name string) bool {
//	return name == "admin"
//}

func main() {
	blacklist := func(name string) bool {
		return name == "Admin"
	}
	registerUser("Admin", blacklist)
	registerUser("root", blacklist)
}
