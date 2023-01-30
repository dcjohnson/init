package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	configFile = "/etc/init.d/init.conf"
)

type Conf struct {
	Shell string `yaml:"shell"`
}

func ParseConfiguration() (Conf, error) {
	c := Conf{}
	byts, err := os.ReadFile(configFile)
	if err != nil {
		return Conf{}, err
	}
	err = yaml.Unmarshal(byts, &c)
	if err != nil {
		return Conf{}, err
	}

	return c, nil
}
