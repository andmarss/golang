package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

const UsernameRegex string = `^(\w){1,18}$`

func main() {
	var usernameInput string

	flag.StringVar(&usernameInput, "username", "Andrew", "Username to check")
	flag.Parse()

	fmt.Println("Username validation checker")
	fmt.Println("Check for username syntax, \"", usernameInput, "\" resulted in: ", CheckUsernameSyntax(usernameInput), "\n")
}

func CheckUsernameSyntax(username string) bool {
	validationResult := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(username)
	return validationResult
}
