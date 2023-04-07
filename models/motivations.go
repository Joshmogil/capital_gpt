package models

import (
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"log"
	"fmt"
)

type Motivation struct {
	Body string
	Policies []string
}

type Motivations struct {
	motivations map[string]Motivation
}

func (motivations *Motivations) loadMotivations(configLocation string) error {
	file, err := ioutil.ReadFile(configLocation)
	if err != nil {
		return err
    }

	data := make(map[string]Motivation)
	err2 := yaml.Unmarshal(file, &data)
	if err2 != nil {
		return err2
	}

	for mName, mStruct := range data {
		fmt.Printf("%s: %s\n  Policies:\n", mName, mStruct.Body)
		for i := range mStruct.Policies {
			fmt.Printf("    %d: %s \n",i ,mStruct.Policies[i])
		}

    }

	motivations.motivations = data
	log.Printf("Successfully loaded motivations from local file %s", configLocation)
	return nil
}