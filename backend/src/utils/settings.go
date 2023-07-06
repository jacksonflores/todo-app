package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	DB string `yaml:"db"`
}

var SettingsStuff Settings

func ParseSettings() {
	data, err := ioutil.ReadFile("../settings.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &SettingsStuff)
	if err != nil {
		log.Fatal(err)
	}
}
