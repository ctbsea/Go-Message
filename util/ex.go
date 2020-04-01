package util

import (
	"gopkg.in/yaml.v2"
	"log"
)

func YamlInterfaceToStruct(str interface{}, result interface{})  {
	yamlString, err := yaml.Marshal(str)
	if err != nil {
		log.Fatal(" Can't Parse")
	}
	unErr := yaml.Unmarshal(yamlString, result)
	if unErr != nil {
		log.Fatal(" Can't Parse")
	}
}
