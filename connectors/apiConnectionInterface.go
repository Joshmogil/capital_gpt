package connectors

type apiConnector interface {
	getResponse(string) string
	testConnection() bool
}