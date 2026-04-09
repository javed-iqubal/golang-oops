package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	response := struct {
		Status     string
		StatusCode int
		Message    string
	}{
		Status:     "OK",
		StatusCode: 200,
		Message:    "Created successfully",
	}

	jsonData, _ := json.Marshal(response)

	fmt.Println(string(jsonData))
}
