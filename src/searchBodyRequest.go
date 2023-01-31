package main

type SearchRequest struct {
	Query  Match    `json:"query"`
	Source []string `json:"_source"`
	Size   int      `json:"size"`
}

type Match struct {
	MatchPhrase Body `json:"match_phrase"`
}

type Body struct {
	Body string `json:"Body"`
}
