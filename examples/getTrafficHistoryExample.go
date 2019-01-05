package main

import (
	"bytes"
	"fmt"
	"github.com/qwer7y/golexaRank"
	"io/ioutil"
	"strings"
)

func main() {
	trafficInfoResponseGroups := "History"
	exampleDomain := "www.github.com"

	// Reading the credentials required for issuing the requests
	accessID := "Add_Your_AccessID_here" // Avoid hardcoding this. Read it from file and add it as an environment variable
	secretAccessKey := "Add_Your_SecretAccessKey_here"

	// Let's see if the GetTrafficInfo function works
	myRange := "31"
	start := "20070801"
	response := golexaRank.GetTrafficHistory(myRange, start, exampleDomain, trafficInfoResponseGroups, accessID, secretAccessKey)
	if err != nil {
		println("Error reading response.Body")
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	newStr := buf.String()
	fmt.Println(newStr)
}
