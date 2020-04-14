package model

type Post struct {
	Company  string `json:"company"`
	Section  string `json:"section"`
	Grade    string `json:"grade"`
	Name     string `json:"name"`
	ItemType int    `json:"itemType"`
}
