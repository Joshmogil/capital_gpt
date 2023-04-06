package state

type motivation struct {
	goal string
	policies []string
}

type Motivations struct {
	motivations map[int]motivation
}