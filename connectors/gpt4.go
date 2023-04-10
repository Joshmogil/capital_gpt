package connectors

type gpt4Connection struct {
	apiToken string
	openaiOrg string
}

func (conn *gpt4Connection) getResponse(string) string {
	return "response"
}

func (conn *gpt4Connection) testConnection() bool {
	return true
}

