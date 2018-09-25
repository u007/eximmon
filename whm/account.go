package whm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Account struct {
	OutgoingMailSuspended int    `json:"outgoing_mail_suspended"`
	OutgoingMailHold      int    `json:"outgoing_mail_hold"`
	IP                    string `json:"ip"`

	UID           string `json:"uid"`
	StartDate     string `json:"startdate"`
	DiskUsed      string `json:"diskused"`
	DistLimit     string `json:"disklimit"`
	User          string `json:"user"`
	Owner         string `json:"owner"`
	Suspended     int    `json:"suspended"`
	SuspendTime   int    `json:"suspendtime"`
	SuspendReason string `json:"suspendreason"`
	IsLocked      int    `json:"is_locked"`
	Plan          string `json:"plan"`
}

func AccountInfo(domain string) (Account, error) {
	Log("Domain: %s", domain)
	urlString := apiURI + "accountsummary?api.version=1&domain=" + url.QueryEscape(domain)
	// Log("Connecting to '%s'", ApiHost)
	conn, err := WHMDialer()
	if err != nil {
		return Account{}, err
	}

	// conn.SetReadDeadline(time.Now().UTC().Add(time.Minute * 5))
	defer conn.Close()
	clientConn := httputil.NewClientConn(conn, nil)

	Log("calling: %s", urlString)
	req, err := http.NewRequest("GET", urlString, nil)
	authStr := fmt.Sprintf("whm %s:%s", ApiUser, ApiToken)
	// Log("authorization: %s", authStr)
	req.Header.Set("Authorization", authStr)

	resp, err := clientConn.Do(req)
	if err != nil {
		return Account{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}
	Log("result %#v", string(body))

	var record ApiResponse
	if err := json.Unmarshal(body, &record); err != nil {
		return Account{}, err
	}

	Log("json %#v", record)
	if record.Metadata.Result == 1 {
		if len(record.Data.Accounts) < 1 {
			return Account{}, fmt.Errorf("Account not found")
		}

		return record.Data.Accounts[0], nil
	} else {
		// Log("metadata: %#v", record.Metadata)
		return Account{}, fmt.Errorf(record.Metadata.Reason)
	}

	return Account{}, fmt.Errorf("json error?")
}
