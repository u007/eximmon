package whm

import "strings"

type CPanelApiResponse struct {
	Func   string          `json:"func"`
	Result CPanelApiResult `json:"result"`
}

type CPanelApiResult struct {
	MetaData MetaData  `json:"metadata"`
	Warnings *[]string `json:"warnings"`
	Errors   *[]string `json:"errors"`
	Messages *[]string `json:"messages"`
	Data     int       `json:"data"`
	Status   int       `json:"status"`
}

func (a CPanelApiResult) ErrorMessage() string {
	if a.Errors == nil {
		return ""
	}

	return strings.Join(*a.Errors, ", ")
}

type ApiResponse struct {
	Metadata MetaData `json:"metadata"`
	Data     Data     `json:"data"`
}

type MetaData struct {
	Version int    `json:"version"`
	Reason  string `json:"reason"`
	Result  int    `json:"result"`
	Command string `json:"command"`
}

type Data struct {
	Domains  []Domain  `json:"domains"`
	Accounts []Account `json:"acct"`
	UserData UserData  `json:"userdata"`

	Reason string `json:"reason"`
	Result string `json:"result"`
	Error  string `json:"error"`
}

var ApiUser = "root"
var ApiToken = ""
var ApiHost = "127.0.0.1"
var apiURI = "/json-api/"
var cpanelApiURL = apiURI + "cpanel"

var Log func(string, ...interface{})

func cPanelApiURL(module string, function string, user string) string {
	return apiURI + "cpanel?api.version=1" + "&cpanel_jsonapi_user=" + user +
		"&cpanel_jsonapi_apiversion=3" +
		"&cpanel_jsonapi_module=" + module +
		"&cpanel_jsonapi_func=" + function
}
