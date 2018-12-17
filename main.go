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
	"os"
	"time"
	"github.com/anaskhan96/soup"
)

/**
This does something
param:
returns:
 */
func create_v4_signature (map[string]string) (string, map[string]string)  {
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

/**
This does something
param:
returns:
 */
func sign (key string, msg string) {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
}

/**
This does something
param:
returns:
 */
func get_signature_key (key string, dateStamp string, regionName  string, serviceName  string)  {

}

/**
This does something
param:
returns:
 */
func urlinfo (domain string, responseGroup string) (url string, headers map[string]string) {
	params := make(map[string]string)
	params["Action"] = "UrlInfo"
	params["Url"] = domain
	params["ResponseGroup"] = responseGroup
	URL, headers := create_v4_signature(params)
	return URL, headers
}

/**
The following function is used to get the traffic history of the given domain
# TODO: Make the myRange and start parameters over-ridable
param:
returns:
 */
func traffichistory (domain string, responseGroup string) (string, map[string]string) {
	myRange := "31"
	start := "20070801"
	params := make(map[string]string)
	params["Action"] = "TrafficHistory"
	params["Url"] = domain
	params["ResponseGroup"] = responseGroup
	params["Range"] = myRange
	params["Start"] = start
	URL, headers := create_v4_signature(params)
	return URL, headers
}

/**
This function provides us the informaiton on sites linking in for a specified domain
param: Domain name
param: Response group
returns:
 */
func siteslinkingin (domain string, responseGroup string) (string, map[string]string) {
	params := make(map[string]string)
	params["Action"] = "SitesLinkingIn"
	params["Url"] = domain
	params["ResponseGroup"] = responseGroup
	URL, headers := create_v4_signature(params)
	return URL, headers
}

/**
This function provides the category browse information for a specified domain
param: Domain name
param: Path TODO: Wtf is this supposed to be?
param: responseGroup
param: descriptions
returns: URL, headers generated from the create_v4_signature function
 */
func cat_browse (domain string, path string, responseGroup string, descriptions string) (string, map[string]string) {
	params := make(map[string]string)
	params["Action"] = "CategoryListings"
	params["ResponseGroup"] = "Listings"
	// Add quote(path) to the below
	params["Path"] = "Listings"
	params["Descriptions"] = descriptions
	URL, headers := create_v4_signature(params)
	return URL, headers
}

/**
This does something
params: A URL and the headers obtained from the create_v4_signature function
returns: A beautiful soup object where the characters are encoded in utf-8 and the object is formatted as 'XML'
 */
func return_output (url string, headers map[string]string)  {
	// Look up CheckRedirect policies and see if one should be added here
	client := &http.Client {}
	request, _ := http.NewRequest("GET", url, nil)
	for index, element := range headers {
		request.Header.Add(index, element)
	}
	response, err := client.Do(request)
	if err != nil {
		os.Exit(1)
	}
	// return response
}

func main() {
	urlInfoResponseGroups := "RelatedLinks,Categories,Rank,ContactInfo,RankByCountry,UsageStats,Speed,Language,OwnedDomains,LinksInCount,SiteData,AdultContent"
	trafficInfoResponseGroups := "History"
	sitesLinkingInResponseGroup := "SitesLinkingIn"
	exampleDomain := "www.github.com"
	// Let's see if the urlInfo function works
	urlinfo(exampleDomain, urlInfoResponseGroups)
	// Let's see if the trafficInfo function works
	traffichistory(exampleDomain, trafficInfoResponseGroups)
	// Let's see if the sitesLinkingIn function works
	siteslinkingin(exampleDomain, sitesLinkingInResponseGroup)
	// Let's see if the cat_browse function works
}
