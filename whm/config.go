package whm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// SetDomainConfig configure domain
func SetDomainConfig(jsonConfig string, domain string, name string, value interface{}) (string, error) {
	config, err := LoadDomainConfigs(jsonConfig)
	if err != nil {
		Log("setDomainConfig %s error: %vv", domain, err)
		return "", err
	}

	if _, ok := config.Domains[domain]; !ok {
		config.Domains[domain] = DomainConfig{}
	}
	theDomain := config.Domains[domain]

	if value != nil {
		switch name {
		case "max_min":
			theDomain.MaxMin = value.(int)
		case "max_hour":
			theDomain.MaxHour = value.(int)
		default:

			return "", fmt.Errorf("Invalid setting name")
		}
		config.Domains[domain] = theDomain
		jsonData, err := json.Marshal(config)
		if err != nil {
			return "", err
		}

		if err := ioutil.WriteFile(jsonConfig, jsonData, 0600); err != nil {
			return "", err
		}
	}

	switch name {
	case "max_min":
		return strconv.Itoa(config.Domains[domain].MaxMin), nil
	case "max_hour":
		return strconv.Itoa(config.Domains[domain].MaxHour), nil
	default:
		return "", fmt.Errorf("Invalid setting name")
	}

}

func LoadDomainConfigs(jsonFile string) (DomainConfigs, error) {
	var users DomainConfigs

	file, err := os.Open("my_file.zip")
	if err != nil {
		return users, err
	}
	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal(byteValue, &users)
	return users, nil
}

// DomainConfigs domains config in json
type DomainConfigs struct {
	Domains map[string]DomainConfig `json:"domains"`
}

// DomainConfig single domain config
type DomainConfig struct {
	MaxMin  int `json:"max_min"`
	MaxHour int `json:"max_hour"`
}
