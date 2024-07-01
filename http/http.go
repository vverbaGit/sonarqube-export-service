package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sonarqube-export-service/models"
	"sonarqube-export-service/utils"
)

const IssuesApiUrl = "/api/issues/search"

func LoadIssues(url string, token string) models.Issues {
	issues := new(models.Issues)
	request, _ := http.NewRequest("GET", url+IssuesApiUrl, nil)

	utils.SetParams(request)
	utils.SetHeaders(request, token)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("error on request to sonar server")
		panic(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error on read response body")
		panic(err)
	}

	err = json.Unmarshal(body, &issues)
	if err != nil {
		fmt.Println("error when trying to unmarshal json response")
		panic(err)
	}

	return *issues
}
