package whm

import (
	"encoding/json"
	"io/ioutil"
)

func SetDomainConfig(jsonConfig string, domain string, name string, value interface{}) (string, error) {
	config := loadDomainConfigs(jsonConfig)

	if val, ok := config.Domains[domain]; ok {
		config.Domains[domain] = DomainConfig{}
	}
	theDomain := config.Domains[domain]

	if value != nil {
		switch name {
		case "max_min":
			theDomain.MaxMin = int(value)
		case "max_hour":
			theDomain.MaxHour = int(value)
		default:

			return nil, error.Error("Invalid setting name")
		}
		config.Domains[domain] = theDomain
		jsonData, err := json.Marshal(config)
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(jsonFile, jsonData, 0600); err != nil {
			return nil, err
		}
	}

	switch name {
	case "max_min":
		return domain.MaxMin, nil
	case "max_hour":
		return domain.MaxHour, nil
	default:
		return nil, error.Error("Invalid setting name")
	}

}

func loadDomainConfigs(jsonFile String) DomainConfigs {
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users DomainConfigs

	json.Unmarshal(byteValue, &users)
	return users
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
