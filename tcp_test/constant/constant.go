package constant

import "errors"

type SendMessage struct {
	MType    MessageType `json:"mtype"`    //消息类型
	SType    int64       `json:"s_type"`   //发送类型
	MUser    int64       `json:"m_user"`   //发送者
	MData    string      `json:"m_data"`   //发送数据
	Receiver int64       `json:"receiver"` //接受者 只有是单播才会存在
}
type MessageType int

// 定义消息的三种状态
const (
	Connection MessageType =  1
	Send MessageType = 2
	Leave MessageType = 3
	Close MessageType  = 4
)

// 定义发送消息的类型
const (
	SingleSendType    int64 = 1
	BroadcastSendType int64 = 2
)

// 定义错误的消息结构体
var (
	JoinChatExists               = errors.New("已经加入过房间")
	SingleChatRevicerError       = errors.New("单播用户不存在")
	SingleChatRevicerAcceptError = errors.New("单播用户上线")
	SendTypeError                = errors.New("发送未有该类型")
	DeleteNotExists              = errors.New("删除不存在数据")
)
