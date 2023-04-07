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
