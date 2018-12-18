package main

import (
	"encoding/hex"
	"testing"
)

func TestCreateV4Signature(t *testing.T) {

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
