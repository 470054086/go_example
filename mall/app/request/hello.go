package request

// index 的请求参数
type IndexRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Sex      int    `json:"sex"`
}

type IndexResponse struct {
	Id uint `json:"id"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	Sex int `json:"sex"`
}
