package models

import (
    "testing"
)

func TestLoadMotivations(t *testing.T) {
    configLocation := "../inputs/motivations.yaml"

    m := &Motivations{}

	err := m.loadMotivations(configLocation)

	if err != nil {
        t.Fatalf(`TestLoadMotivations = %v,`, err)
    }
}

func TestLoadResources(t *testing.T) {
    configLocation := "../inputs/resources.yaml"

    m := &Resources{}

	err := m.loadResources(configLocation)

	if err != nil {
        t.Fatalf(`TestLoadResources = %v,`, err)
    }
}