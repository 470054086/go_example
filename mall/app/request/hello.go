package request

// index 的请求参数
type IndexRequest struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
	Sex      int    `json:"sex" binding:"required"`
}

type IndexResponse struct {
	Id       uint   `json:"id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Sex      int    `json:"sex"`
}

// list 请求参数
type ListRequest struct {
	Mobile string `json:"mobile"`
	Sex    int    `json:"sex"`
}
type ListResponse struct {
	List     []*UserListResponse `json:"list"`
	Total    int                 `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
}
type UserListResponse struct {
	Mobile  string
	Sex     int
	SexName string
}

// 登录数据请求
type LoginRequest struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}
