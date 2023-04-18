package ai

type AiConnector interface {
	getResponse(string) string
}