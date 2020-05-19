package SendMobile

import "context"

type SendMobile interface {
	// 发送短信
	Send(mobile *SendMessage) error
	Star(context.Context, context.CancelFunc)
	recover()
	Stop()
}