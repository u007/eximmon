package whm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func SuspendEmail(email string) error {
	Log("Suspending %s", email)

	domain := email[strings.Index(email, "@")+1:]
	info, err := UserDataInfo(domain)
	if err != nil {
		Log("UserDataInfo error: %v", err)
		return err
	}

	urlString := cPanelApiURL("Email", "suspend_outgoing", info.User) + "&email=" + url.QueryEscape(email)
	Log("calling: %s", urlString)
	time.Sleep(1000 * time.Millisecond)
	// Log("Connecting to '%s'", ApiHost)
	conn, err := WHMDialer()
	if err != nil {
		return err
	}
	// conn.SetReadDeadline(time.Now().UTC().Add(time.Minute * 5))
	defer conn.Close()
	clientConn := httputil.NewClientConn(conn, nil)

	req, err := http.NewRequest("GET", urlString, nil)
	req.Header.Set("Authorization", fmt.Sprintf("whm %s:%s", ApiUser, ApiToken))

	resp, err := clientConn.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	Log("Reading body")
	body, err := ioutil.ReadAll(resp.Body)
	Log("called: %s", body)
	if err != nil {
		return err
	}
	Log("result %#v", string(body))

	var record CPanelApiResponse
	if err := json.Unmarshal(body, &record); err != nil {
		return err
	}

	// Log("json %#v", record)
	if record.Result.Status == 1 {
		return nil
	} else {
		Log("metadata: %#v", record.Result.MetaData)
		return fmt.Errorf(record.Result.ErrorMessage())
	}
}

func UnSuspendEmail(email string) error {
	Log("UnSuspending: %s", email)

	domain := email[strings.Index(email, "@")+1:]
	info, err := UserDataInfo(domain)
	if err != nil {
		return err
	}

	urlString := cPanelApiURL("Email", "unsuspend_outgoing", info.User) + "&email=" + url.QueryEscape(email)
	// Log("Connecting to '%s'", ApiHost)
	conn, err := WHMDialer()
	if err != nil {
		return err
	}
	// conn.SetReadDeadline(time.Now().UTC().Add(time.Minute * 5))
	defer conn.Close()
	clientConn := httputil.NewClientConn(conn, nil)

	// Log("calling: %s", urlString)
	req, err := http.NewRequest("GET", urlString, nil)
	req.Header.Set("Authorization", fmt.Sprintf("whm %s:%s", ApiUser, ApiToken))

	resp, err := clientConn.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	Log("result %#v", string(body))

	var record CPanelApiResponse
	if err := json.Unmarshal(body, &record); err != nil {
		return err
	}

	// Log("json %#v", record)
	if record.Result.Status == 1 {
		return nil
	} else {
		Log("metadata: %#v", record.Result.MetaData)
		return fmt.Errorf(record.Result.ErrorMessage())
	}
}

//================

func SuspendAccountByEmail(email string) error {
	Log("Suspending: %s", email)

	domain := email[strings.Index(email, "@")+1:]
	info, err := UserDataInfo(domain)
	if err != nil {
		return err
	}

	urlString := apiURI + "suspend_outgoing_email?api.version=1&user=" + url.QueryEscape(info.User)
	// Log("Connecting to '%s'", ApiHost)
	conn, err := WHMDialer()
	if err != nil {
		return err
	}
	defer conn.Close()
	clientConn := httputil.NewClientConn(conn, nil)

	Log("calling: %s", urlString)
	req, err := http.NewRequest("GET", urlString, nil)
	req.Header.Set("Authorization", fmt.Sprintf("whm %s:%s", ApiUser, ApiToken))

	resp, err := clientConn.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// body2, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// Log("result %#v", string(body2))
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var records []ArxivRecord
	// err = json.Unmarshal(body, &records)
	// if err != nil {
	// 	panic(fmt.Errorf("json error: %v", err.Error()))
	// }
	dec := json.NewDecoder(resp.Body)
	for dec.More() {
		var record ApiResponse
		// decode an array value (Message)
		if err := dec.Decode(&record); err != nil {
			return err
		}

		Log("json %#v", record)
		if record.Metadata.Result == 1 {
			return nil
		} else {
			Log("metadata: %#v", record.Metadata)
			return fmt.Errorf(record.Metadata.Reason)
		}
	} //has json

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("json error? %#v", body)
}

//https://documentation.cpanel.net/display/DD/WHM+API+1+Functions+-+suspend_outgoing_email
func UnSuspendAccountByEmail(email string) error {
	Log("Suspending: %s", email)

	domain := email[strings.Index(email, "@")+1:]
	info, err := UserDataInfo(domain)
	if err != nil {
		return err
	}

	urlString := apiURI + "suspend_outgoing_email?api.version=1&user=" + url.QueryEscape(info.User)
	// Log("Connecting to '%s'", ApiHost)
	conn, err := WHMDialer()
	if err != nil {
		return err
	}
	defer conn.Close()
	clientConn := httputil.NewClientConn(conn, nil)

	Log("calling: %s", urlString)
	req, err := http.NewRequest("GET", urlString, nil)
	req.Header.Set("Authorization", fmt.Sprintf("whm %s:%s", ApiUser, ApiToken))

	resp, err := clientConn.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// body2, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// Log("result %#v", string(body2))
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var records []ArxivRecord
	// err = json.Unmarshal(body, &records)
	// if err != nil {
	// 	panic(fmt.Errorf("json error: %v", err.Error()))
	// }
	dec := json.NewDecoder(resp.Body)
	for dec.More() {
		var record ApiResponse
		// decode an array value (Message)
		if err := dec.Decode(&record); err != nil {
			return err
		}

		Log("json %#v", record)
		if record.Metadata.Result == 1 {
			return nil
		} else {
			Log("metadata: %#v", record.Metadata)
			return fmt.Errorf(record.Metadata.Reason)
		}
	} //has json

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("json error? %#v", body)
}

func UnsuspendAccountByEmail(email string) error {
	Log("Suspending: %s", email)

	domain := email[strings.Index(email, "@")+1:]
	info, err := UserDataInfo(domain)
	if err != nil {
		return err
	}

	urlString := apiURI + "unsuspend_outgoing_email?api.version=1&user=" + url.QueryEscape(info.User)
	// Log("Connecting to '%s'", ApiHost)
	conn, err := WHMDialer()
	if err != nil {
		return err
	}
	defer conn.Close()
	clientConn := httputil.NewClientConn(conn, nil)

	Log("calling: %s", urlString)
	req, err := http.NewRequest("GET", urlString, nil)
	req.Header.Set("Authorization", fmt.Sprintf("whm %s:%s", ApiUser, ApiToken))

	resp, err := clientConn.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// body2, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// Log("result %#v", string(body2))
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var records []ArxivRecord
	// err = json.Unmarshal(body, &records)
	// if err != nil {
	// 	panic(fmt.Errorf("json error: %v", err.Error()))
	// }
	dec := json.NewDecoder(resp.Body)
	for dec.More() {
		var record ApiResponse
		// decode an array value (Message)
		if err := dec.Decode(&record); err != nil {
			return err
		}

		Log("json %#v", record)
		if record.Metadata.Result == 1 {
			return nil
		} else {
			Log("metadata: %#v", record.Metadata)
			return fmt.Errorf(record.Metadata.Reason)
		}
	} //has json

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("json error? %#v", body)
}
