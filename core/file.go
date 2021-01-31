package core

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Conf config yaml file format
type Conf struct {
	Hosts []string `yaml:"hosts"`
}

// Read read yaml file
func read(location string) []string {
	yamlFile, err := ioutil.ReadFile(location)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Failed to read config file: %s", location)
	}

	c := Conf{}
	err = yaml.Unmarshal(yamlFile, &c)

	if err != nil {
		log.Fatalf("Failed to read config file: %s", err.Error())
	}
	return c.Hosts
}
