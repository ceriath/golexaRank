package main

import (
	"fmt"
	"github.com/qwer7y/golexaRank"
	"io/ioutil"
	"strings"
)

func main() {
	// Change this to an array of strings
	// Check for the Python urlencode equivalent of this
	categoryBrowseInfoResponseGroup := "Categories%2CRelatedCategories%2CLanguageCategories%2CLetterBars"
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

	// Let's see if the GetCategoryBrowseInformation function works
	// TODO: Change this below
	path := "True"
	description := "True"
	response := golexaRank.GetCategoryBrowseInformation(exampleDomain, path, categoryBrowseInfoResponseGroup, description, accessID, secretAccessKey)
	if response.StatusCode != 200 {
		println("Response status code: " + response.Status)
		println("Response headers: ")
		for key, value := range response.Header {
			print(key + ": ")
			for v := range value {
				println(v)
			}
		}
	}
}