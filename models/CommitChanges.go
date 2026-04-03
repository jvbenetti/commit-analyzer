package models

type CommitChanges struct {
	SHA     string `json:"sha"`
	Changes int    `json:"changes"`
}
