package actions

import (
	"management/enums"
)

type Response struct {
	Message    string            `json:"message,omitempty"`
	Type       enums.MessageType `json:"type,omitempty"`
	Data       interface{}       `json:"data,omitempty"`
	Errors     interface{}       `json:"errors,omitempty"`
	Pagination interface{}       `json:"pagination,omitempty"`
}
