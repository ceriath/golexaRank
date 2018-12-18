package main

import (
	"testing"
)

func TestCreateV4Signature(t *testing.T) {

}

func TestSign(t *testing.T) {
	key := []byte("Sucker")
	dateStamp := []byte("20120215")
	testDrive := sign(key, dateStamp)
	print(testDrive)
}

func TestGetSignatureKey(t *testing.T) {

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

func TestName(t *testing.T) {

}
