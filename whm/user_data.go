package whm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type UserData struct {
	IP string `json:"ip"`

	User         string `json:"user"`
	Group        string `json:"group"`
	Owner        string `json:"owner"`
	DocumentRoot string `json:"documentroot"`
	HomeDir      string `json:"homedir"`
}

func UserDataInfo(domain string) (UserData, error) {
	// Log("Domain: %s", domain)
	urlString := apiURI + "domainuserdata?api.version=1&domain=" + url.QueryEscape(domain)
	// Log("Connecting to '%s'", ApiHost)
	conn, err := WHMDialer()
	if err != nil {
		return UserData{}, err
	}
	defer conn.Close()
	clientConn := httputil.NewClientConn(conn, nil)

	// Log("calling: %s", urlString)
	req, err := http.NewRequest("GET", urlString, nil)
	authStr := fmt.Sprintf("whm %s:%s", ApiUser, ApiToken)
	// Log("authorization: %s", authStr)
	req.Header.Set("Authorization", authStr)

	resp, err := clientConn.Do(req)
	if err != nil {
		return UserData{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UserData{}, err
	}
	// Log("result %#v", string(body))

	var record ApiResponse
	if err := json.Unmarshal(body, &record); err != nil {
		return UserData{}, err
	}

	// Log("json %#v", record)
	if record.Metadata.Result == 1 {
		return record.Data.UserData, nil
	} else {
		// Log("metadata: %#v", record.Metadata)
		return UserData{}, fmt.Errorf(record.Metadata.Reason)
	}

	// return UserData{}, fmt.Errorf("json error?")
}
