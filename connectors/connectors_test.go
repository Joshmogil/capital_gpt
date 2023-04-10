package connectors

import (
    "testing"
)

func TestGpt4Connection(t *testing.T) {
    apiToken := "../inputs/motivations.yaml"

    m := &Motivations{}

	err := m.loadMotivations(configLocation)

	if err != nil {
        t.Fatalf(`TestLoadMotivations = %v,`, err)
    }
}