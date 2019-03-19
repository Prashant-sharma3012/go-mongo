package responses

import "encoding/json"

type ErrorMessage struct {
	errorMessage string `json:"errorMessage"`
	statusCode   int32  `json:"statusCode"`
}

func (e *ErrorMessage) Error() string {
	str, _ := json.Marshal(e)

	return string(str)
}

func ServerError(message string) *ErrorMessage {
	err := &ErrorMessage{}
	err.statusCode = 501
	err.errorMessage = message

	return err
}

func ClientError(message string) *ErrorMessage {
	err := &ErrorMessage{}
	err.statusCode = 401
	err.errorMessage = message

	return err
}

func NotFound(message string) *ErrorMessage {
	err := &ErrorMessage{}
	err.statusCode = 404
	err.errorMessage = message

	return err
}

func NotAuthorized(message string) *ErrorMessage {
	err := &ErrorMessage{}
	err.statusCode = 403
	err.errorMessage = message

	return err
}
