package main

import (
	"bytes"
	"fmt"
	"github.com/qwer7y/golexaRank"
	"io/ioutil"
	"strings"
)

func main() {
	// Choose the response groups you want
	urlInfoResponseGroupList := []string{"RelatedLinks", "Categories", "Rank", "ContactInfo", "RankByCountry", "UsageStats", "Speed", "Language", "OwnedDomains", "LinksInCount", "SiteData", "AdultContent"}
	var urlInfoResponseGroups string
	length := len(urlInfoResponseGroupList)
	for i, val := range urlInfoResponseGroupList {
		if i != (length - 1) {
			urlInfoResponseGroups += val
			urlInfoResponseGroups += "%2C"
		} else {
			urlInfoResponseGroups += val
		}
	}
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
	if err != nil {
		println("Error reading response.Body")
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	newStr := buf.String()
	fmt.Println(newStr)
}
