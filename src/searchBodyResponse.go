package main

type SearchResponse struct {
	Hits HitsParent `json:"hits"`
}

type HitsParent struct {
	Hits []HitsChild `json:"hits"`
}

type HitsChild struct {
	ID     string `json:"_id"`
	Source Source `json:"_source"`
}

type Source struct {
	Body    string `json:"Body"`
	Date    string `json:"Date"`
	From    string `json:"From"`
	Subject string `json:"Subject"`
	To      string `json:"To"`
}
