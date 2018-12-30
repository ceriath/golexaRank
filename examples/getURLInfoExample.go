package main

import (
	"bytes"
	"fmt"
	"github.com/qwer7y/golexaRank"
	"io/ioutil"
	"strings"
)

func main() {
	// Change this to an array of strings
	// Check for the Python urlencode equivalent of this
	urlInfoResponseGroups := "RelatedLinks%2CCategories%2CRank%2CContactInfo%2CRankByCountry%2CUsageStats%2CSpeed%2CLanguage%2C" +
		"OwnedDomains%2CLinksInCount%2CSiteData%2CAdultContent"
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

	// Let's see if the GetUrlInfo function works
	response := golexaRank.GetUrlInfo(exampleDomain, urlInfoResponseGroups, accessID, secretAccessKey)
	if response.StatusCode != 200 {
		println("Response status code: " + response.Status)
		println("Response headers: ")
		//for key, value := range response.Header {
		//	print(key + ": ")
		//	for v := range value {
		//		println(v)
		//	}
		//}
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(response.Body)
		if err != nil {
			println("Problem reading response.Body")
		}
		newStr := buf.String()
		fmt.Printf(newStr)
	} else {
		println("Success!")
	}
}
