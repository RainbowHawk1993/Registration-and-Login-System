package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	var choice int
	fmt.Println("1: Register\n2: Login")
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 2 {
		login()
	}
	if choice == 1 {
		registration()
		main()
	}
}

func registration() {
	var username string
	var password string
	fmt.Print("Enter username (without spaces):")
	fmt.Scan(&username)
	fmt.Print("Enter password:")
	fmt.Scan(&password)

	txt := ".txt"
	var filename string
	filename += username
	filename += txt
	//fmt.Println("File/Username: ", filename)
	//fmt.Println("password: ", password)

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		// file does not exist, so we create a new file. This helps us avoid truncating the old one
		txtFile, err := os.Create(filename)
		if err != nil {
			exit("Failed to create a txt file.")
		}

		defer txtFile.Close()

		_, err2 := txtFile.WriteString(username)
		fmt.Fprintf(txtFile, "\n%s\n", password)

		if err2 != nil {
			exit("Failed to write to txt file.")
		}

		fmt.Println("Succesfully registered")
		return
	} else {
		// file exists
		fmt.Println("This username is already taken")
		return
	}
}

func login() {
	var username, password string
	fmt.Print("Enter username:")
	fmt.Scan(&username)
	fmt.Print("Enter password:")
	fmt.Scan(&password)

	txt := ".txt"
	var filename string
	filename += username
	filename += txt

	file, err := os.Open(filename)
	if err != nil {
		exit("Failed to open txt file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	if username == text[0] && password == text[1] {
		fmt.Print("Succesfully logged in!\n")
		os.Exit(1)
	} else {
		fmt.Print("There was an error logging in\n")
		main()
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
