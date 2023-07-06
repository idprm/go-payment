package rest_errors

import "net/http"

type RestErr struct {
	Error        bool        `json:"error"`
	StatusCode   int         `json:"status_code"`
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"error_message"`
}

func NewBadRequestError() *RestErr {
	return &RestErr{
		Error:      true,
		StatusCode: http.StatusBadRequest,
		Message:    "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Error:      true,
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func NewInternalServerError() *RestErr {
	return &RestErr{
		Error:      true,
		StatusCode: http.StatusInternalServerError,
		Message:    "internal_server_error",
	}
}

func NewBadGatewayError() *RestErr {
	return &RestErr{
		Error:      true,
		StatusCode: http.StatusBadGateway,
		Message:    "bad_gateway",
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Error:      true,
		StatusCode: http.StatusForbidden,
		Message:    message,
	}
}

func NewValidateError(message interface{}) *RestErr {
	return &RestErr{
		Error:        true,
		StatusCode:   http.StatusForbidden,
		ErrorMessage: message,
	}
}
