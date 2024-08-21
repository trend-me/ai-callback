package exceptions

const (
	ValidationErrorCode                         = 1
	UnknownErrorCode                            = 2
	AiFactoryErrorCode                          = 3
	QueueErrorCode                              = 4
	PromptRoadMapErrorCode                      = 5
	GetPromptRoadMapConfigErrorCode             = 6
	PayloadValidatorNotFoundErrorCode           = 7
	PayloadValidatorErrorCode                   = 8
	UpdatePromptRoadMapConfigExecutionErrorCode = 9
)

func NewValidationError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     true,
		Notify:    true,
		ErrorType: "Validation Error",
		message:   messages,
		Code:      ValidationErrorCode,
	}
}

func NewUnknownError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    false,
		ErrorType: "Unknown Error",
		message:   messages,
		Code:      UnknownErrorCode,
	}
}

func NewQueueError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    true,
		ErrorType: "Queue Error",
		message:   messages,
		Code:      QueueErrorCode,
	}
}

func NewPromptRoadMapNotFoundError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     true,
		Notify:    true,
		ErrorType: "Prompt Road Map Not Found Error",
		message:   messages,
		Code:      PromptRoadMapErrorCode,
	}
}

func NewGetPromptRoadMapConfigError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    true,
		ErrorType: "Get Prompt Road Map Config Error",
		message:   messages,
		Code:      GetPromptRoadMapConfigErrorCode,
	}
}

func NewPayloadValidatorNotFoundError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    true,
		ErrorType: "Prompt Road Map Config Error",
		message:   messages,
		Code:      PayloadValidatorNotFoundErrorCode,
	}
}

func NewPayloadValidatorError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    true,
		ErrorType: "Payload Validator Error",
		message:   messages,
		Code:      PayloadValidatorErrorCode,
	}
}

func NewUpdatePromptRoadMapConfigExecutionError(messages ...string) ErrorType {
	return ErrorType{
		Abort:     false,
		Notify:    true,
		ErrorType: "Update PromptRoadMapConfigExecution Error",
		message:   messages,
		Code:      UpdatePromptRoadMapConfigExecutionErrorCode,
	}
}
