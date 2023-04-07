package models

//maybe store the state in pinecone like auto-gpt?

type State struct {
	Resources Resources
	Motivations Motivations
	PreviousOutcomes PreviousOutcomes
}