package handler

import "fmt"

func CheckError(err error, message string) {
	if err != nil {
		fmt.Printf("Message: %s\nError: %s", message, err.Error())
	}
}
