package commom
/**
	参数错误类型
 */
type ParamsError struct {
	message string
	Err error
}
func (p ParamsError) Error() string  {
	return p.message
}

func NewParamsError(err error,message string) ParamsError {
	return ParamsError{
		message: message,
		Err:   err,
	}
}
