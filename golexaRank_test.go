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

// Test for internal helper functions end

// Tests for exported functions start
func TestGetUrlInfo(t *testing.T) {

}

func TestGetTrafficHistory(t *testing.T) {

}

func TestGetSitesLinkingIn(t *testing.T) {

}

func TestGetCategoryBrowseInformation(t *testing.T) {

}

func TestReturnOutput(t *testing.T) {

}

// Tests for exported functions end
