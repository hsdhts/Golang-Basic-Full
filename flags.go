package main

import (
	"flag"
	"fmt"
)

func main() {

	var host *string = flag.String("host", "localhost", "put your database localhost")
	var user *string = flag.String("user", "root", "put your username database")
	var password *string = flag.String("password", "root", "put your password database")

	flag.Parse()

	fmt.Println("Host: ", *host)
	fmt.Println("user :", *user)
	fmt.Println("Password :", *password)
}
