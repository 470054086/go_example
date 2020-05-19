package constaent

type LoginError int

// 定义数字
const (
	UserNotExistError LoginError = 1000 //用户不存在
	UserPasswordError LoginError = 1001 //用户不存在或者密码错误
)
// 定义map
var d = map[LoginError]string{
	UserPasswordError: "用户或者密码错误",
	UserNotExistError: "用户或者密码错误",
}

func (u LoginError) String() string {
	i := d[u]
	return i
}
