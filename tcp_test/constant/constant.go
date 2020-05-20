package constant

type SendMessage struct {
	MType MessageType `json:"mtype"`
	MUser int64       `json:"m_user"`
	MData interface{} `json:"m_data"`
}
type MessageType int

// 定义消息的三种状态
const (
	Send MessageType = iota
	Leave
	Close
)
