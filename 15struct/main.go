package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}
	return a / b, nil
}

func login(username, password string) (string, error) {
	if username != "admin" || password != "admin" {
		return "", errors.New("login failed")
	}
	return "login success", nil
}

func main() {

	result, err := divide(20, 10)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result, err)

	status, errorCode := login("admin", "admin")

	if errorCode != nil {
		fmt.Println(errorCode)
		return
	}
	fmt.Println(status)
}
