package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Need to be moved to a configuration file
const ZINC_INDEX_NAME = "/email_index"
const ZINC_SEARCH_URL = "http://localhost:4080/es" + ZINC_INDEX_NAME
const ZINCSEARCH_USER = "admin"
const ZINCSEARCH_PASSWORD = "admin"

func searchBody(term string) []EmailData {
	requestBody := SearchRequest{
		Query: Match{
			MatchPhrase: Body{
				Body: term,
			},
		},
		Source: []string{"_id", "Date", "Subject", "From", "To", "Body"},
		Size:   1000,
	}

	jsonRequest, errParsing := json.Marshal(requestBody)
	handleError(errParsing)

	request, errRequest := http.NewRequest(
		"POST",
		ZINC_SEARCH_URL+"/_search",
		bytes.NewBuffer(jsonRequest))
	handleError(errRequest)
	request.SetBasicAuth(ZINCSEARCH_USER, ZINCSEARCH_PASSWORD)

	client := http.Client{}
	response, errResponse := client.Do(request)
	handleError(errResponse)

	defer func(Body io.ReadCloser) {
		errOnClose := Body.Close()
		handleError(errOnClose)
	}(response.Body)

	responseData := SearchResponse{}

	errDecoding := json.NewDecoder(response.Body).Decode(&responseData)
	handleError(errDecoding)

	var emailDataList []EmailData

	dataList := responseData.Hits.Hits
	for i := 0; i < len(dataList); i++ {
		dataSource := dataList[i].Source
		emailData := EmailData{
			Id:      dataList[i].ID,
			Subject: strings.TrimSpace(dataSource.Subject),
			From:    strings.TrimSpace(dataSource.From),
			To:      strings.TrimSpace(dataSource.To),
			Body:    strings.TrimSpace(dataSource.Body),
			Date:    strings.TrimSpace(dataSource.Date),
		}
		emailDataList = append(emailDataList, emailData)
	}

	return emailDataList
}
