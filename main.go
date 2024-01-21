package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	url := "https://footy-fixtures-backend.onrender.com/fixtures/"
	var response *http.Response

	// Using os.Args to check if command is run with a specific date
	// If no date is provided then pass the date command was run
	if len(os.Args) == 1 {
		var err error

		// Get current date and convert it to a string to pass
		// Into the URL
		time := time.Now()
		timeString := strings.SplitAfter(time.String(), " ")

		response, err = http.Get(url + strings.Trim(timeString[0], " "))
		if err != nil {
			fmt.Println("Error making request:", err)
		}
	} else {
		var err error

		response, err = http.Get(url + os.Args[1])
		if err != nil {
			fmt.Println("Error making request:", err)
		}
	}

	defer response.Body.Close()

	// Need to adjust to make a little more elegant
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}

	fmt.Println(string(body))

}
