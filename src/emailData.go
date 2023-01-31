package main

// EmailData Type of object to return to the client
type EmailData struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
	From    string `json:"from"`
	To      string `json:"to"`
	Body    string `json:"body"`
	Date    string `json:"date"`
}
