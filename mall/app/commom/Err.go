package commom

/**
参数错误类型
*/
type ParamsError struct {
	message string
	code    int
	Err     error
}

func (p ParamsError) Error() string {
	return p.message
}
func (p ParamsError) Code() int {
	return p.code
}

func NewParamsError(err error, message string) ParamsError {
	return ParamsError{
		message: message,
		Err:     err,
	}
}

func NewParamsErrorCode(err error, message string, code int) ParamsError {
	return ParamsError{
		message: message,
		Err:     err,
		code:    code,
	}
}
