package enums

import (
	"errors"
	"strings"
)

type MessageType string

const (
	Error   MessageType = "error"
	Success MessageType = "success"
	Warning MessageType = "warning"
)

var (
	invalidType = errors.New("Inalid type")
)

func (mt *MessageType) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	messageType := MessageType(strings.Trim(string(b), `"`))
	switch messageType {
	case Error, Success, Warning:
		*mt = messageType
		return nil
	}
	return invalidType
}

func (mt MessageType) IsValid() bool {
	switch mt {
	case Error, Success, Warning:
		return true
	}
	return false
}
