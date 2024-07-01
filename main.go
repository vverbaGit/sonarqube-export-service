package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"sonarqube-export-service/config"
	"sonarqube-export-service/http"
	"sonarqube-export-service/models"
	"sonarqube-export-service/utils"
)

func SaveToClipboard(text string) error {
	err := clipboard.WriteAll(text)
	if err != nil {
		return fmt.Errorf("failed to write to clipboard: %v", err)
	}
	return nil
}

func main() {
	conf := config.Load()

	issues := http.LoadIssues(conf.BaseUrl, conf.Token)

	issueMap := make(map[models.IssueKey]map[string]int)

	for _, issue := range issues.Issues {
		if issueMap[issue.GetIssueKey()] == nil {
			issueMap[issue.GetIssueKey()] = make(map[string]int)
		}
		issueMap[issue.GetIssueKey()][issue.Message] += 1
	}

	types := utils.GetIssuesTypes(issueMap)

	result := "\n"
	for _, key := range types {
		result += fmt.Sprintf("%s %s:\n", key.Type, key.Severity)
		for k, v := range issueMap[key] {
			result += fmt.Sprintf("%v - x%v times \n", k, v)
		}
		result += fmt.Sprintf("\n")
	}

	fmt.Println(result)

	err := SaveToClipboard(result)
	if err != nil {
		panic(err)
	}
}
