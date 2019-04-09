package lib

import "time"

// RawReq 表示原生请求的结构。
type RawReq struct {
	ID  int64
	Req []byte
}
// RawResp 表示原生响应的结构。
type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}
// RetCode 表示结果代码的类型。
type RetCode int
// CallResult 表示调用结果的结构。
type CallResult struct {
	ID     int64         // ID。
	Req    RawReq        // 原生请求。
	Resp   RawResp       // 原生响应。
	Code   RetCode       // 响应代码。
	Msg    string        // 结果成因的简述。
	Elapse time.Duration // 耗时。
}

// 载荷发生器的接口，用于面向接口编程，初始化载荷发生器，对载荷发生器进行操作
type Generator interface {
	Start() bool
	Stop() bool
	Status() uint32
	//	获取调用计数，每次启动会重启这个计数
	CallCount() int64
}
