package models

import "fmt"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func BadRequestError(message string) Error {
	return Error{
		Code:    "BadRequest",
		Message: message,
	}
}

func ExistsError() Error {
	return Error{
		Code: "Exists",
	}
}

func ServerError(err error) Error {
	return Error{
		Code:    "ServerError",
		Message: fmt.Sprintf("%v", err),
	}
}

func NotFoundError() Error {
	return Error{
		Code: "NotFound",
	}
}
