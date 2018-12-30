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
	fileReadBytes, err := ioutil.ReadFile("credentials.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileReadString := string(fileReadBytes)
	fileReadStringSplit := strings.Split(fileReadString, "\n")
	accessID := fileReadStringSplit[0]
	secretAccessKey := fileReadStringSplit[1]

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
