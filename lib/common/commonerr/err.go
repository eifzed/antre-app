package commonerr

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	ErrorList []*ErrorFormat
	Code      int
}

type ErrorFormat struct {
	ErrorName        string `json:"error_name"`
	ErrorDescription string `json:"error_description"`
}

func NewError(errorCode int, errorName, errorDesc string) *ErrorMessage {
	return &ErrorMessage{
		Code: errorCode,
		ErrorList: []*ErrorFormat{
			{
				ErrorName:        errorName,
				ErrorDescription: errorDesc,
			},
		},
	}
}

func ErrorBadRequest(errorName, errorDesc string) *ErrorMessage {
	return NewError(http.StatusBadRequest, errorName, errorDesc)
}

func ErrorAlreadyExist(errorName, errorDesc string) *ErrorMessage {
	return NewError(http.StatusConflict, errorName, errorDesc)
}

func ErrorUnauthorized(errorDesc string) *ErrorMessage {
	return NewError(http.StatusUnauthorized, "Unauthorized", errorDesc)
}

func (errorMessage *ErrorMessage) Error() string {
	return errorMessage.ToString()
}

func (errorMessage *ErrorMessage) ToString() string {
	b, _ := json.Marshal(errorMessage)
	return string(b)
}

func (errorMessage *ErrorMessage) GetCode() int {
	return errorMessage.Code
}

func ErrorNotFound(context string) *ErrorMessage {
	return NewError(http.StatusNotFound, context, context+" not found")
}
