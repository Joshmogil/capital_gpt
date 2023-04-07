package models

type resource struct {
	name string
	resourceType string
	resourceAttributes map[string]string
	policies []string
}

type Resources struct {
	motivations map[int]resource
}