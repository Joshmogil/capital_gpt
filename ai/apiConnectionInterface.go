package ai

type aiConnector interface {
	getResponse(string) string
	testConnection() bool
}