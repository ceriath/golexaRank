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
	"crypto/hmac"
	"github.com/anaskhan96/soup"
	"hash"
	"net/http"
	"net/url"
	"os"
	"time"
)

/**
This does something
param:
returns:
 */
func createV4Signature(map[string]string) (string, map[string]string)  {
	method := "GET"
	service := "awis"
	host := "awis.us-west-1.amazonaws.com"
	region := "us-west-1"
	endpoint := "https://awis.amazonaws.com/api"
	requestParameters := "awis"

	// We need to create a date for headers and the credential string
	t := time.Now().UTC()
	amzDate := t.Format("%Y%m%dT%H%M%SZ")
	dateStamp := t.Format("%Y%m%d")

	// Now to create a canonical request
	canonicalUri := "/api"
	canonicalQuerystring := requestParameters
	canonicalHeaders := "host:" + host + "\n" + "x-amz-date:" + amzDate + "\n"
	signedHeaders := "host;x-amz-date"
	payloadHash := 0 // TODO
	canonicalRequest := method + "\n" + canonicalUri + "\n" + canonicalQuerystring + "\n" + canonicalHeaders + "\n" + signedHeaders + "\n" + payloadHash

	// Create string to sign
	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := dateStamp + "/" + region + "/" + service + "/" + "aws4_request"
	stringToSign := algorithm + "\n" + amzDate + "\n" + credentialScope + "\n" +  hashlib.sha256(canonicalRequest.encode('utf8')).hexdigest() // TODO

	// Calculate signature
	// TODO
	accessID := "TODO"
	secretAccessKey := "TODO"
	// TODO
	signingKey := getSignatureKey(secretAccessKey, dateStamp, region, service)

	// Sign the stringToSign using the signingKey TODO
	signature := hmac.new(signingKey, (stringToSign).encode('utf-8'), hashlib.sha256).hexdigest()

	// Add signing information to the request
	authorizationHeader := algorithm + " " + "Credential=" + accessId + "/" + credentialScope + ", " +  "SignedHeaders=" + signedHeaders + ", " + "Signature=" + signature
	headers := make(map[string]string)
	headers["X-Amz-Date"] = amzDate
	headers["Authorization"] = authorizationHeader
	headers["Content-Type"] = "application/xml"
	headers["Accept"] = "application/xml"

	// Create request url
	requestUrl := endpoint + "?" + canonicalQuerystring

	return requestUrl, headers
}


/**
TODO: https://stackoverflow.com/questions/21961615/why-doesnt-go-allow-nested-function-declarations-functions-inside-functions
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
func getSignatureKey(key string, dateStamp string, regionName  string, serviceName  string)  {

}

/**
This function provides us the URL information for a given domain
param: Domain name of the site
param: responseGroup for the GetUrlInfo function
returns: The response with the URL information as a http.Response type
 */
func GetUrlInfo(domainURL string, responseGroup string) *http.Response {
	params := make(map[string]string)
	params["Action"] = "UrlInfo"
	params["Url"] = domainURL
	params["ResponseGroup"] = responseGroup
	URL, headers := createV4Signature(params)
	return ReturnOutput(URL, headers)
}

/**
This function provides us the traffic history of the given domain
# TODO: Make the myRange and start parameters override-able
param: Domain name of the site
param: ResponseGroup for getting the traffic history
returns: The response with the traffic history data as a http.Response type
 */
func GetTrafficHistory(domainURL string, responseGroup string) *http.Response {
	myRange := "31"
	start := "20070801"
	params := make(map[string]string)
	params["Action"] = "TrafficHistory"
	params["Url"] = domainURL
	params["ResponseGroup"] = responseGroup
	params["Range"] = myRange
	params["Start"] = start
	URL, headers := createV4Signature(params)
	return ReturnOutput(URL, headers)
}

/**
This function provides us the information on sites linking in for a specified domain
param: Domain name of the site
param: Response group
returns: The response with the get sites linking data as a http.Response type
 */
func GetSitesLinkingIn(domainURL string, responseGroup string) *http.Response {
	params := make(map[string]string)
	params["Action"] = "SitesLinkingIn"
	params["Url"] = domainURL
	params["ResponseGroup"] = responseGroup
	URL, headers := createV4Signature(params)
	return ReturnOutput(URL, headers)
}

/**
This function provides the category browse information for a specified domain
param: Domain name
param: Path TODO: Wtf is this supposed to be?
param: responseGroup
param: descriptions
returns: URL, headers generated from the createV4Signature function
 */
func GetCategoryBrowseInformation(domainURL string, path string, responseGroup string, descriptions string) *http.Response {
	params := make(map[string]string)
	params["Action"] = "CategoryListings"
	params["ResponseGroup"] = "Listings"
	// Add quote(path) to the below
	params["Path"] = "Listings"
	params["Descriptions"] = descriptions
	URL, headers := createV4Signature(params)
	return ReturnOutput(URL, headers)
}

/**
This function takes in a domain name, headers for the request and returns an http.Response type
param: Domain name string
param: A map with headers
returns: An HTTP response type object
 */
func ReturnOutput(domainURL string, headers map[string]string) *http.Response  {
	// Look up CheckRedirect policies and see if one should be added here
	client := &http.Client {}
	request, _ := http.NewRequest("GET", domainURL, nil)
	for index, element := range headers {
		request.Header.Add(index, element)
	}
	response, err := client.Do(request)
	if err != nil {
		os.Exit(1)
	}
	return response
}

/**
This function takes in the http Response type and parses into usable XML
 */
func httpResponseToXML() {

}

func main() {
	urlInfoResponseGroups := "RelatedLinks,Categories,Rank,ContactInfo,RankByCountry,UsageStats,Speed,Language,OwnedDomains,LinksInCount,SiteData,AdultContent"
	trafficInfoResponseGroups := "History"
	sitesLinkingInResponseGroup := "SitesLinkingIn"
	categoryBrowseInfoResponseGroup := "Categories,RelatedCategories,LanguageCategories,LetterBars"
	exampleDomain := "www.github.com"
	// Let's see if the GetUrlInfo function works
	GetUrlInfo(exampleDomain, urlInfoResponseGroups)
	// Let's see if the trafficInfo function works
	GetTrafficHistory(exampleDomain, trafficInfoResponseGroups)
	// Let's see if the sitesLinkingIn function works
	GetSitesLinkingIn(exampleDomain, sitesLinkingInResponseGroup)
	// Let's see if the GetCategoryBrowseInformation function works
	// TODO: Change this below
	path := "True"
	description := "True"
	GetCategoryBrowseInformation(exampleDomain, path, categoryBrowseInfoResponseGroup, description)
}
