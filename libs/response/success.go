package response

import (
	type_response "github.com/uni-school/api-gateaway/libs/response/types"
)

func ResponseSuccess[T any](message string, data T) type_response.ResponseData[T] {
	return type_response.ResponseData[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}
