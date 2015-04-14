package gohadoopxml

import (
	"encoding/xml"
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
