package client

import "context"

type CallFunc func(ctx context.Context, ServiceMethod string, arg interface{}, reply interface{}) error
type GoFunc func(ctx context.Context, ServiceMethod string, arg interface{}, reply interface{}, done chan *Call) *Call

// 定义Client Wrapper
type Wrapper interface {
	WrapCall(option *SGOption, callFunc CallFunc) CallFunc
	WrapGo(option *SGOption, goFunc GoFunc) GoFunc
}