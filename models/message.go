package models

// Message model for responses
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
