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
	accessID := "Add_Your_AccessID_here" // Avoid hardcoding this. Read it from file and add it as an environment variable
	secretAccessKey := "Add_Your_SecretAccessKey_here"

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
