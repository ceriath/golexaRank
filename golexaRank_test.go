package golexaRank

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

// Tests for internal helper functions begin
func TestCreateV4Signature(t *testing.T) {
	parameterSet := make(map[string]string)
	parameterSet["Action"] = "CategoryListings"
	parameterSet["Url"] = "www.github.com"
	parameterSet["ResponseGroup"] = "History"
	parameterSet["Range"] = "31"
	parameterSet["Start"] = "20070801"

	// Reading the required credentials
	fileReadBytes, err := ioutil.ReadFile("credentials.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileReadString := string(fileReadBytes)
	fileReadStringSplit := strings.Split(fileReadString, "\n")
	accessID := fileReadStringSplit[0]
	secretAccessKey := fileReadStringSplit[1]

	requestUrl, headers := createV4Signature(parameterSet, accessID, secretAccessKey)
	if !strings.Contains(requestUrl, "Action") {
		t.Error("Issue with request")
	}
	if headers == nil {
		t.Error("Issue with headers")
	}
}

func TestSign(t *testing.T) {
	key := []byte("Some fucking test case string")
	dateStamp := []byte("20120215")
	testDrive := sign(key, dateStamp)
	calculatedValue := hex.EncodeToString(testDrive)
	correctValue := "f2f077be4009e87c842f68c51f42600352d63ea696e97be5e8a2aef51fa5168b"
	if calculatedValue != correctValue {
		t.Error("")
	}
}

func TestGetSignatureKey(t *testing.T) {
	key := "Some fucking key string"
	ds := "Some stupid datestamp"
	regionName := "Some idiotic region name"
	serviceName := "Blah blah blah"
	calculatedValue := getSignatureKey(key, ds, regionName, serviceName)
	calculatedHexValue := hex.EncodeToString(calculatedValue)
	correctValue := "d66455cc65b63e63c7efb9341602ff784d8f68c3f36ecaac4a783ba1d2ddc280"
	if calculatedHexValue != correctValue {
		t.Error("")
	}
}

func TestReturnOutput(t *testing.T) {
	params := make(map[string]string)
	params["Action"] = "UrlInfo"
	params["ResponseGroup"] = "RelatedLinks%2CCategories%2CRank%2CContactInfo%2CRankByCountry%2CUsageStats%2CSpeed%2CLanguage%2C" +
		"OwnedDomains%2CLinksInCount%2CSiteData%2CAdultContent"
	params["Url"] = "www.github.com"

	// Reading the required credentials
	fileReadBytes, err := ioutil.ReadFile("credentials.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileReadString := string(fileReadBytes)
	fileReadStringSplit := strings.Split(fileReadString, "\n")
	accessID := fileReadStringSplit[0]
	secretAccessKey := fileReadStringSplit[1]

	domainURL, headers := createV4Signature(params, accessID, secretAccessKey)
	sampleReturnOutput := returnOutput(domainURL, headers)
	if sampleReturnOutput.StatusCode != 200 {
		println("Error!")
		print("Response status code: ")
		println(sampleReturnOutput.StatusCode)
	}
}

// Test for internal helper functions end

// Tests for exported functions start
func TestGetUrlInfo(t *testing.T) {
	urlInfoResponseGroupList := []string{"RelatedLinks", "Categories", "Rank", "ContactInfo", "RankByCountry"}
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
	response := GetUrlInfo(exampleDomain, urlInfoResponseGroups, accessID, secretAccessKey)
	if response.StatusCode != 200 {
		t.Error("Status code is not 200!")
	}
}

func TestGetTrafficHistory(t *testing.T) {
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

	// Let's see if the trafficInfo function works
	myRange := "31"
	start := "20070801"
	response := GetTrafficHistory(myRange, start, exampleDomain, trafficInfoResponseGroups, accessID, secretAccessKey)
	if response.StatusCode != 200 {
		println("Status code is not 200")
	}
}

func TestGetSitesLinkingIn(t *testing.T) {

}

func TestGetCategoryBrowseInformation(t *testing.T) {

}

// Tests for exported functions end
