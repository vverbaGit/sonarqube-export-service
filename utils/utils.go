package utils

import (
	"net/http"
	"sonarqube-export-service/models"
	"sort"
)

// SEVERITY
const (
	BlockerSeverity  string = "BLOCKER"
	CriticalSeverity string = "CRITICAL"
	HighSeverity     string = "HIGH"
	MajorSeverity    string = "MAJOR"
	MinorSeverity    string = "MINOR"
	InfoSeverity     string = "INFO"
	LowSeverity      string = "LOW"
)

// TYPE
const (
	CodeSmellType     string = "CODE_SMELL"
	BugType           string = "BUG"
	VulnerabilityType string = "VULNERABILITY"
)

var weightMap = map[models.IssueKey]int{
	{Severity: BlockerSeverity, Type: BugType}:            10,
	{Severity: CriticalSeverity, Type: BugType}:           10,
	{Severity: HighSeverity, Type: BugType}:               9,
	{Severity: MajorSeverity, Type: BugType}:              9,
	{Severity: MinorSeverity, Type: BugType}:              8,
	{Severity: InfoSeverity, Type: BugType}:               8,
	{Severity: LowSeverity, Type: BugType}:                8,
	{Severity: BlockerSeverity, Type: VulnerabilityType}:  7,
	{Severity: CriticalSeverity, Type: VulnerabilityType}: 7,
	{Severity: HighSeverity, Type: VulnerabilityType}:     6,
	{Severity: MajorSeverity, Type: VulnerabilityType}:    6,
	{Severity: MinorSeverity, Type: VulnerabilityType}:    5,
	{Severity: InfoSeverity, Type: VulnerabilityType}:     5,
	{Severity: LowSeverity, Type: VulnerabilityType}:      5,
	{Severity: BlockerSeverity, Type: CodeSmellType}:      4,
	{Severity: CriticalSeverity, Type: CodeSmellType}:     4,
	{Severity: HighSeverity, Type: CodeSmellType}:         3,
	{Severity: MajorSeverity, Type: CodeSmellType}:        3,
	{Severity: MinorSeverity, Type: CodeSmellType}:        2,
	{Severity: InfoSeverity, Type: CodeSmellType}:         2,
	{Severity: LowSeverity, Type: CodeSmellType}:          2,
}

func SetHeaders(request *http.Request, token string) {
	request.Header.Set("Authorization", "Bearer "+token)
}

func SetParams(request *http.Request) {
	query := request.URL.Query()
	query.Add("assignees", "__me__")
	request.URL.RawQuery = query.Encode()
}

func GetIssuesTypes(issueMap map[models.IssueKey]map[string]int) []models.IssueKey {
	types := make([]models.IssueKey, 0, len(issueMap))

	for key := range issueMap {
		types = append(types, key)
	}
	sort.Slice(types, func(i, j int) bool {
		return weightMap[types[i]] > weightMap[types[j]]
	})
	return types
}
