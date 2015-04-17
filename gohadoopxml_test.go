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
