package config

import (
	"encoding/json"
	"io/ioutil"
)

var Conf *Config

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &Conf); err != nil {
		return nil, err
	}
	return Conf, nil
}
