package app

import "time"

type response struct {
	Error  bool `json:"error"`
	Result any  `json:"result"`
}

type reqCreateUpdateTask struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}
