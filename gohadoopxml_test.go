package gohadoopxml

import (
	"testing"
)

func isin(key string, c Configuration) int {
	for i, p := range c.Properties {
		if p.Name == key {
			return i
		}
	}
	return -1
}

func TestMergeConfiguration(t *testing.T) {
	var config1 = Configuration{
		Properties: []Property{
			{
				Name:  "foo",
				Value: "bar",
			},
		},
	}
	var config2 = Configuration{
		Properties: []Property{
			{
				Name:  "fizz",
				Value: "buzz",
			},
		},
	}
	result := MergeConfigurations(config1, config2)
	if i := isin("foo", result); i < 0 {
		t.Fail()
	}
	if i := isin("fizz", result); i < 0 {
		t.Fail()
	}
}

func TestGetPropertyValue(t *testing.T) {
	var config = Configuration{
		Properties: []Property{
			{
				Name:  "foo",
				Value: "bar",
			},
		},
	}
	result, err := GetPropertyValue("foo", config)
	if err != nil || result != "bar" {
		t.Fail()
	}
	result, err = GetPropertyValue("fizz", config)
	if err == nil {
		t.Fail()
	}
}
