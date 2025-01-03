package model

type Blog struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	Timestamp string `json:"timestamp"`
}
