package helpers

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 返回成功的数据
func Success(data interface{}, message string) *Response {

	return &Response{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

func Error(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
