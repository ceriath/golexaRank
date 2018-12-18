package main

import (
	"encoding/hex"
	"testing"
)

func TestCreateV4Signature(t *testing.T) {

}

func TestSign(t *testing.T) {
	key := []byte("Sucker")
	dateStamp := []byte("20120215")
	testDrive := sign(key, dateStamp)
	correctValue := "2dd290c65716bf297ee9feb426fc708002c6ace2dace50b925e93dee5a9141c9"
	if hex.EncodeToString(testDrive) != correctValue {
		t.Error("The hash values don't match")
	}
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
