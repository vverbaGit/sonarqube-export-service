package models

type TextRange struct {
	StartLine   int `json:"startLine"`
	EndLine     int `json:"endLine"`
	StartOffset int `json:"startOffset"`
	EndOffset   int `json:"endOffset"`
}

type Flow struct {
	// Define the fields for Flow if needed
}

type Impact struct {
	SoftwareQuality string `json:"softwareQuality"`
	Severity        string `json:"severity"`
}

type Issue struct {
	Key                        string        `json:"key"`
	Rule                       string        `json:"rule"`
	Severity                   string        `json:"severity"`
	Component                  string        `json:"component"`
	Project                    string        `json:"project"`
	Line                       int           `json:"line"`
	Hash                       string        `json:"hash"`
	TextRange                  TextRange     `json:"textRange"`
	Flows                      []Flow        `json:"flows"`
	Status                     string        `json:"status"`
	Message                    string        `json:"message"`
	Effort                     string        `json:"effort"`
	Debt                       string        `json:"debt"`
	Assignee                   string        `json:"assignee"`
	Author                     string        `json:"author"`
	Tags                       []string      `json:"tags"`
	Type                       string        `json:"type"`
	Scope                      string        `json:"scope"`
	QuickFixAvailable          bool          `json:"quickFixAvailable"`
	MessageFormattings         []interface{} `json:"messageFormattings"`
	CodeVariants               []interface{} `json:"codeVariants"`
	CleanCodeAttribute         string        `json:"cleanCodeAttribute"`
	CleanCodeAttributeCategory string        `json:"cleanCodeAttributeCategory"`
	Impacts                    []Impact      `json:"impacts"`
	IssueStatus                string        `json:"issueStatus"`
}

type Component struct {
	Key       string `json:"key"`
	Enabled   bool   `json:"enabled"`
	Qualifier string `json:"qualifier"`
	Name      string `json:"name"`
	LongName  string `json:"longName"`
	Path      string `json:"path,omitempty"`
}

type Paging struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}

type Issues struct {
	Total       int           `json:"total"`
	P           int           `json:"p"`
	Ps          int           `json:"ps"`
	Paging      Paging        `json:"paging"`
	EffortTotal int           `json:"effortTotal"`
	Issues      []Issue       `json:"issues"`
	Components  []Component   `json:"components"`
	Facets      []interface{} `json:"facets"`
}

type IssueKey struct {
	Severity string `json:"severity"`
	Type     string `json:"type"`
}

func (issue Issue) ToString() string {
	return issue.Message
}
func (issue Issue) GetIssueKey() IssueKey {
	return IssueKey{issue.Severity, issue.Type}
}
