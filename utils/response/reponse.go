package response

type Sucess struct {
	Code   int
	Status bool
	Data   interface{}
}

type Error struct {
	Code   int
	Status bool
	Error  string
}

type SuccessMsg struct {
	Code    int
	Status  bool
	Data    interface{}
	Message string
}

func SuccessResponse(data interface{}) Sucess {
	result := Sucess{
		Code:   200,
		Status: true,
		Data:   data,
	}
	return result
}

func SuccessResponseMsg(data interface{}, msg string) SuccessMsg {
	result := SuccessMsg{
		Code:    200,
		Status:  true,
		Data:    data,
		Message: msg,
	}
	return result
}

func ErrorResponse(code int, err error) Error {
	result := Error{
		Code:   code,
		Status: true,
		Error:  err.Error(),
	}
	return result
}
