package models

import (
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"log"
	"fmt"
	"strings"
    "unicode/utf8"
)

type Resource struct {
	Name string `yaml:"name"`
	ResourceType string `yaml:"resourceType"`
	ResourceAttributes map[string]string `yaml:"resourceAttributes"`
	Policies []string `yaml:"policies"`
}

type Resources struct {
	resources map[string]Resource
}

func (resources *Resources) loadResources(configLocation string) error {
	file, err := ioutil.ReadFile(configLocation)
	if err != nil {
		return err
    }

	data := make(map[string]Resource)
	err2 := yaml.Unmarshal(file, &data)
	if err2 != nil {
		return err2
	}

	for mName, mStruct := range data {
		fmt.Printf("%s:\n  type: %s\n  name: %s\n  Policies:\n", mName, mStruct.ResourceType, mStruct.Name)
		for i := range mStruct.Policies {
			fmt.Printf("    %d: %s \n",i ,mStruct.Policies[i])
		}
		fmt.Printf("  ResourceAttributes:\n")
		for k, v := range mStruct.ResourceAttributes { 
			starredSecret := strings.Repeat("*", utf8.RuneCountInString(v))
			fmt.Printf("    %s: %s\n", k, starredSecret)
		}

    }

	resources.resources = data
	log.Printf("Successfully loaded resources from local file %s", configLocation)
	return nil
}