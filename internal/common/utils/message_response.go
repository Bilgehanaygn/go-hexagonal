package utils

import (
	"github.com/google/uuid"
)

type MessageType string

const (
    MessageTypeError   MessageType = "ERROR"
    MessageTypeInfo    MessageType = "INFO"
    MessageTypeWarning MessageType = "WARNING"
)

type MessageResponse struct {
    Message   string       `json:"message"`
    Type      MessageType  `json:"type"`
    Arguments []string     `json:"arguments,omitempty"`
    ID        *uuid.UUID   `json:"id,omitempty"`
}

func (r *MessageResponse) HasError() bool {
    return r.Type == MessageTypeError
}

func DefaultSuccessResult() MessageResponse {
    return MessageResponse{
        Message: "",
        Type:    MessageTypeInfo,
    }
}

func DefaultErrorResult() MessageResponse {
    return MessageResponse{
        Message: "",
        Type:    MessageTypeError,
    }
}
