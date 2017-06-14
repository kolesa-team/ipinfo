// Package ipinfo provides info on IP address location
// using the http://ipinfo.io service.
package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ipinfoURI = "https://ipinfo.io"

// IPInfo wraps json response
type IPInfo struct {
	IP       string            `json:"ip"`
	Hostname string            `json:"hostname"`
	City     string            `json:"city"`
	Region   string            `json:"region"`
	Country  string            `json:"country"`
	Loc      string            `json:"loc"`
	Org      string            `json:"org"`
	Postal   string            `json:"postal"`
	Company  map[string]string `json:"company"`
}

// MyIP provides information about the public IP address of the client.
func MyIP(token ...string) (*IPInfo, error) {
	return getInfo(fmt.Sprintf("%s/json", ipinfoURI), token)
}

// ForeignIP provides information about the given IP address (IPv4 or IPv6)
func ForeignIP(ip string, token ...string) (*IPInfo, error) {
	return getInfo(fmt.Sprintf("%s/%s/json", ipinfoURI, ip), token)
}

// Undercover code that makes the real call to the webservice
func getInfo(url string, token []string) (*IPInfo, error) {
	tokenValue := ""
	if len(token) == 1 {
		tokenValue = fmt.Sprintf("?token=%s", token[0])
	}

	response, err := http.Get(url + tokenValue)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var ipinfo IPInfo
	err = json.NewDecoder(response.Body).Decode(&ipinfo)
	if err != nil {
		return nil, err
	}

	return &ipinfo, nil
}
