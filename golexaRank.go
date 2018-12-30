/**
TODO:
- Test coverage for the exported functions
- Check the ResponseGroup strings and look for an urlencode equivalent
- More modularity for all functions
- Look into the path and description variable requirements in the GetCategoryBrowseInformation function
- Design the GetTrafficHistory function for modularity. Make the myRange and start parameters override-able
- Review GetCategoryBrowseInformation and check if it requires a domainURL
- Format response body string to XML
- Documentify the examples for ease of use
*/

package golexaRank

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"sort"
	"time"
)

/**
This function creates the HTTP request URL and corresponding headers for the request in accordance to AWS's Signature
Version 4 Signing Process
params: A map with the required parameters for the request i.e Action, Url, ResponseGroups, etc
returns: The URL string and a map with the headers for the request
*/
func createV4Signature(requestParams map[string]string, accessID string, secretAccessKey string) (string, map[string]string) {
	// An initial set of variables for the request
	method := "GET"
	service := "awis"
	host := "awis.us-west-1.amazonaws.com"
	region := "us-west-1"
	endpoint := "https://awis.amazonaws.com/api"

	// Creating the canonical query string using the request parameters
	requestParameters := ""
	sortedKeySet := make([]string, 0, len(requestParams))
	for key := range requestParams {
		sortedKeySet = append(sortedKeySet, key)
	}
	sort.Strings(sortedKeySet)
	for _, key := range sortedKeySet {
		requestParameters += key
		requestParameters += "="
		requestParameters += requestParams[key]
		requestParameters += "&"
	}

	// Eliminating the last '&'
	requestParameters = requestParameters[:len(requestParameters)-1]
	canonicalQuerystring := requestParameters

	// Creating a date which will be used for the headers and credential string
	t := time.Now().UTC()
	amzDate := t.Format("20060102T150405Z")
	dateStamp := t.Format("20060102")

	// On to creating a canonical request
	canonicalUri := "/api"
	canonicalHeaders := "host:" + host + "\n" + "x-amz-date:" + amzDate + "\n"
	signedHeaders := "host;x-amz-date"
	payloadHashCreator := sha256.New()
	payloadHashCreator.Write([]byte(""))
	payloadHash := hex.EncodeToString(payloadHashCreator.Sum(nil))
	canonicalRequest := method + "\n" + canonicalUri + "\n" + canonicalQuerystring + "\n" + canonicalHeaders + "\n" + signedHeaders + "\n" + payloadHash

	// Creating a string to sign
	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := dateStamp + "/" + region + "/" + service + "/" + "aws4_request"
	canonicalRequestHashCreator := sha256.New()
	canonicalRequestHashCreator.Write([]byte(canonicalRequest))
	canonicalRequestHash := hex.EncodeToString(canonicalRequestHashCreator.Sum(nil))
	stringToSign := algorithm + "\n" + amzDate + "\n" + credentialScope + "\n" + canonicalRequestHash

	// Calculating the signature of the string
	signingKey := getSignatureKey(secretAccessKey, dateStamp, region, service)
	signature := hex.EncodeToString(sign(signingKey, []byte(stringToSign)))

	// Adding the signing information to the request
	authorizationHeader := algorithm + " " + "Credential=" + accessID + "/" + credentialScope + ", " + "SignedHeaders=" + signedHeaders + ", " + "Signature=" + signature

	// Creating the headers for the request
	headers := make(map[string]string)
	headers["Accept"] = "application/xml"
	headers["Authorization"] = authorizationHeader
	headers["Content-Type"] = "application/xml"
	headers["X-Amz-Date"] = amzDate

	// Creating a request url string
	requestUrl := endpoint + "?" + canonicalQuerystring
	return requestUrl, headers
}

/**
This function provides the SHA256 cryptographic hash of the input message with the input key
More on this at: https://docs.aws.amazon.com/general/latest/gr/signature-v4-examples.html
params: The key and message as byte arrays
returns: The SHA256 cryptographic hash of the message with the given key as a byte array
*/
func sign(key []byte, msg []byte) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write(msg)
	return mac.Sum(nil)
}

/**
This function takes in a key, the dateStamp, AWS region name, and AWS service name to create a signature key that follows
AWS's request format. The calculated value is returned as a byte array.
params: A key string, dateStamp string, regionName string, serviceName string
returns: A byte array of the signature
*/
func getSignatureKey(key string, dateStamp string, regionName string, serviceName string) []byte {
	kDate := sign([]byte("AWS4"+key), []byte(dateStamp))
	kRegion := sign(kDate, []byte(regionName))
	kService := sign(kRegion, []byte(serviceName))
	kSigning := sign(kService, []byte("aws4_request"))
	return kSigning
}

/**
This function issues the custom request we've created and returns the response as an http.Response type
params: A request URL string and a map with the request headers
returns: The response to the HTTP request as an http.Response type
*/
func returnOutput(requestURL string, headers map[string]string) *http.Response {
	// Look up CheckRedirect policies and see if one should be added here
	client := &http.Client{}
	request, _ := http.NewRequest("GET", requestURL, nil)
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
This function provides us the URL information for a given domain
params: The domain name string, a responseGroup string, and API credentials for the GetUrlInfo function
returns: The response with the URL information as an http.Response type
*/
func GetUrlInfo(domainURL string, responseGroup string, accessID string, secretAccessKey string) *http.Response {
	params := make(map[string]string)
	params["Action"] = "UrlInfo"
	params["ResponseGroup"] = responseGroup
	params["Url"] = domainURL
	URL, headers := createV4Signature(params, accessID, secretAccessKey)
	return returnOutput(URL, headers)
}

/**
This function provides us the traffic history of the given domain
params: Domain name of the site, a responseGroup string for getting the traffic history, and the API credentials
returns: The response with the traffic history data as an http.Response type
*/
func GetTrafficHistory(domainURL string, responseGroup string, accessID string, secretAccessKey string) *http.Response {
	myRange := "31"
	start := "20070801"
	params := make(map[string]string)
	params["Action"] = "TrafficHistory"
	params["Range"] = myRange
	params["ResponseGroup"] = responseGroup
	params["Start"] = start
	params["Url"] = domainURL
	URL, headers := createV4Signature(params, accessID, secretAccessKey)
	return returnOutput(URL, headers)
}

/**
This function provides us with the information on the sites linking in to a specified domain
params: A domain URL string, a response group string with the required response groups, and the API credentials
returns: The response with the data of the sites linking into a we data as an http.Response type
*/
func GetSitesLinkingIn(domainURL string, responseGroup string, accessID string, secretAccessKey string) *http.Response {
	params := make(map[string]string)
	params["Action"] = "SitesLinkingIn"
	params["ResponseGroup"] = responseGroup
	params["Url"] = domainURL
	URL, headers := createV4Signature(params, accessID, secretAccessKey)
	return returnOutput(URL, headers)
}

/**
This function provides the category browse information for a specified domain
params: A domain name string, a path string, a responseGroup string, a descriptions string, and the API credentials
returns: Category Browse information of the given domain as an http.Response type
*/
func GetCategoryBrowseInformation(domainURL string, path string, responseGroup string, descriptions string, accessID string, secretAccessKey string) *http.Response {
	params := make(map[string]string)
	params["Action"] = "CategoryListings"
	params["Descriptions"] = descriptions
	// Add quote(path) to the below
	params["Path"] = "Listings"
	params["ResponseGroup"] = "Listings"
	URL, headers := createV4Signature(params, accessID, secretAccessKey)
	return returnOutput(URL, headers)
}
