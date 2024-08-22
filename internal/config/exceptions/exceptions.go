package exceptions

const (
	ValidationErrorCode                         = 1
	UnknownErrorCode                            = 2
	AiFactoryErrorCode                          = 3
	QueueErrorCode                              = 4
	PromptRoadMapNotFoundErrorCode              = 5
	GetPromptRoadMapConfigErrorCode             = 6
	UpdatePromptRoadMapConfigExecutionErrorCode = 9
)

func NewValidationError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     true,
		Notify:    true,
		ErrorType: "Validation Error",
		Message:   messages,
		Code:      ValidationErrorCode,
	}
}

func NewUnknownError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    false,
		ErrorType: "Unknown Error",
		Message:   messages,
		Code:      UnknownErrorCode,
	}
}

func NewQueueError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    true,
		ErrorType: "Queue Error",
		Message:   messages,
		Code:      QueueErrorCode,
	}
}

func NewPromptRoadMapNotFoundError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     true,
		Notify:    true,
		ErrorType: "Prompt Road Map Not Found Error",
		Message:   messages,
		Code:      PromptRoadMapNotFoundErrorCode,
	}
}

func NewGetPromptRoadMapConfigError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    true,
		ErrorType: "Get Prompt Road Map Config Error",
		Message:   messages,
		Code:      GetPromptRoadMapConfigErrorCode,
	}
}
func NewUpdatePromptRoadMapConfigExecutionError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    true,
		ErrorType: "Update PromptRoadMapConfigExecution Error",
		Message:   messages,
		Code:      UpdatePromptRoadMapConfigExecutionErrorCode,
	}
}
