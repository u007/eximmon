package whm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Domain struct {
	DocRoot            string `json:"docroot"`
	Domain             string `json:"domain"`
	DomainType         string `json:"domain_type"`
	IPV4               string `json:"ipv4"`
	IP4VSSL            string `json:"ipv4_ssl"`
	IPV6               string `json:"ipv6"`
	IPV6IsDedicated    int    `json:"ipv6_is_dedicated"`
	ModSecurityEnabled int    `json:"modsecurity_enabled"`
	ParentDomain       string `json:"parent_domain"`
	PHPVersion         string `json:"php_version"`
	Port               string `json:"port"`
	PortSSL            string `json:"port_ssl"`
	User               string `json:"user"`
	UserOwner          string `json:"user_owner"`
}

func Domains() ([]Domain, error) {
	urlString := apiURI + "get_domain_info?api.version=1"
	// Log("Connecting to '%s'", ApiHost)
	conn, err := WHMDialer()
	if err != nil {
		return []Domain{}, err
	}
	defer conn.Close()
	clientConn := httputil.NewClientConn(conn, nil)

	Log("calling: %s", urlString)
	req, err := http.NewRequest("GET", urlString, nil)
	authStr := fmt.Sprintf("whm %s:%s", ApiUser, ApiToken)
	// Log("authorization: %s", authStr)
	req.Header.Set("Authorization", authStr)

	resp, err := clientConn.Do(req)
	if err != nil {
		return []Domain{}, err
	}

	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	for dec.More() {
		var record ApiResponse
		// decode an array value (Message)
		if err := dec.Decode(&record); err != nil {
			return []Domain{}, err
		}

		Log("json %#v", record)
		if record.Metadata.Result == 1 {
			return record.Data.Domains, nil
		} else {
			Log("metadata: %#v", record.Metadata)
			return []Domain{}, fmt.Errorf(record.Metadata.Reason)
		}
	} //has json

	return []Domain{}, fmt.Errorf("json error?")
}
