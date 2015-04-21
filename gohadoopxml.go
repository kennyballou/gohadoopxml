package gohadoopxml

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var Version string

type Property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

type Configuration struct {
	XMLName    xml.Name   `xml:"configuration"`
	Properties []Property `xml:"property"`
}

func ParseXML(filename string) (Configuration, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		log.Println("Error occurred while opening xml file")
		return Configuration{}, err
	}
	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var config Configuration
	xml.Unmarshal(xmlData, &config)

	return config, nil
}

func GetPropertyValue(key string, config Configuration) (string, error) {
	for _, p := range config.Properties {
		if key == p.Name {
			return p.Value, nil
		}
	}
	return "", errors.New("Key not found")
}

func MergeConfigurations(configs ...Configuration) Configuration {
	var new_config Configuration
	for _, config := range configs {
		new_config.Properties = append(new_config.Properties,
			config.Properties...)
	}
	return new_config
}
