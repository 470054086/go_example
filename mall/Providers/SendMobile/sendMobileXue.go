package SendMobile

import (
	"context"
	"fmt"
	"time"
)

type SendMobileXue struct {
	SendChanCount int
	sendChan      chan *SendMessage
	ctx           context.Context
	cancel        context.CancelFunc
}

// 创造一个发送端
var G_MobileSend *SendMobileXue

func NewSendXue(number int) *SendMobileXue {
	G_MobileSend = &SendMobileXue{
		SendChanCount: number,
		sendChan:      make(chan *SendMessage, number),
	}
	return G_MobileSend
}

func (s *SendMobileXue) Star(ctx context.Context, cancelFunc context.CancelFunc) {
	s.ctx = ctx
	s.cancel = cancelFunc
	go s.recover()
}

func (s *SendMobileXue) Send(mobile *SendMessage) error {
	s.sendChan <- mobile
	return nil
}

func (s *SendMobileXue) recover() {

	timer := time.NewTicker(time.Second * 1)
	for {
		select {
		case v := <-s.sendChan:
			go func() {
				fmt.Printf("发送短信手机号为%s,消息内容为%s", v.Mobile, v.Msg)
				fmt.Println()
			}()

		case <-timer.C:

		case <-s.ctx.Done():
			s.close()
			goto Ent
		}
	}
	Ent:
		fmt.Println("我执行完了")
}

func (s *SendMobileXue) Stop() {
	s.cancel()
}

func (s *SendMobileXue) close() {
	fmt.Printf("我的发送函数进行停止了")
	s.sendChan = nil
	s.SendChanCount = 0
	G_MobileSend = nil
}
