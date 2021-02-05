package response

import (
	"fmt"
)

type AppError struct {
	Inner      error  `json:"inner"`
	Msg        string `json:"msg"`
	StackTrace string `json:"-"`
}

func WrapError(err error, msgF string, msgArgs ...interface{}) AppError {
	return AppError{
		Inner: err,
		Msg:   fmt.Sprintf(msgF, msgArgs...),
		// StackTrace: string(debug.Stack()),
	}
}

func (err AppError) Error() string {
	return err.Msg
}
