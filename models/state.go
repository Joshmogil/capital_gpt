package models

//maybe store the state in pinecone like auto-gpt?

type State struct {
	resources Resources
	motivations Motivations
	previousOutcomes Outcome
}