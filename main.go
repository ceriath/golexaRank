// Basically we need to provide 3 input variables:
// 1. Access_Key_ID
// 2. Secret_Access_Key
// 3. Website
// And get the output for the given website

/*
Just tinkering around with Go
This is based off of https://github.com/ashim888/awis
*/

package main

import (
	"crypto/sha256"
	"fmt"
	"crypto"
	"crypto/hmac"
	"hash"
	"net/http"
	"net/url"
	"time"
	"github.com/anaskhan96/soup"
)

func create_v4_signature (map[string]string)  {
	method := "GET"
	service := "awis"
	host := "awis.us-west-1.amazonaws.com"
	region := "us-west-1"
	endpoint := "https://awis.amazonaws.com/api"
	request_parameters := "awis"

	// We need to create a date for headers and the credential string
	t := time.Now().UTC()
	amzdate := t.Format("%Y%m%dT%H%M%SZ")
	datestamp := t.Format("%Y%m%d")

	// Now to create a canonical request
	canonical_uri := "/api"
	request_parameters := 0 // TODO
	canonical_querystring := request_parameters
	canonical_headers := "host:" + host + "\n" + "x-amz-date:" + amzdate + "\n"
	signed_headers := "host;x-amz-date"
	payload_hash := 0 // TODO
	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signed_headers + "\n" + payload_hash

	// Create string to sign
	algorithm := "AWS4-HMAC-SHA256"
	credential_scope := datestamp + "/" + region + "/" + service + "/" + "aws4_request"
	string_to_sign := algorithm + "\n" +  amzdate + "\n" +  credential_scope + "\n" +  hashlib.sha256(canonical_request.encode('utf8')).hexdigest() // TODO

	// Calculate signature
	// TODO
	access_id := "TODO"
	secret_access_key := "TODO"
	// TODO
	signing_key := get_signature_key(secret_access_key, datestamp, region, service)

	// Sign the string_to_sign using the signing_key TODO
	signature := hmac.new(signing_key, (string_to_sign).encode('utf-8'), hashlib.sha256).hexdigest()

	// Add signing information to the request
	authorization_header := algorithm + " " + "Credential=" + access_id + "/" + credential_scope + ", " +  "SignedHeaders=" + signed_headers + ", " + "Signature=" + signature
	headers := make(map[string]string)
	headers["X-Amz-Date"] = amzdate
	headers["Authorization"] = authorization_header
	headers["Content-Type"] = "application/xml"
	headers["Accept"] = "application/xml"

	// Create request url
	request_url := endpoint + "?" + canonical_querystring

	return request_url, headers
}

func sign (key string, msg string) {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
}

func get_signature_key (key string, dateStamp string, regionName  string, serviceName  string)  {

}

func urlinfo (domain string, response_group map[string]string) {
	params := make(map[string]string)
	params["Action"] = "UrlInfo"
	params["Url"] = domain
	params["ResponseGroup"] = response_group
}

func main() {
	fmt.Println("Hello world")
	TRAFFICINFO_RESPONSE_GROUPS := "History"
	SITESLINKINGIN_RESPONSE_GROUP := "SitesLinkingIn"
}
