package models

import (
	"github.com/trend-me/ai-callback/internal/config/exceptions"
)

type Request struct {
	PromptRoadMapConfigName        string
	PromptRoadMapStep              int
	PromptRoadMapConfigExecutionId string
	OutputQueue                    string
	Model                          string
	Error                          *exceptions.ErrorType
	Metadata                       map[string]any
}
