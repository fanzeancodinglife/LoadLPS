package lib

import (
	"errors"
	"fmt"
)

// GoTickets 表示Goroutine票池的接口。
type GoTickets interface {
	// 拿走一张票。
	Take()
	// 归还一张票。
	Return()
	// 票池是否已被激活。
	Active() bool
	// 票的总数。
	Total() uint32
	// 剩余的票数。
	Remainder() uint32
}

// 具体接口的实现
type myGoTickets struct {
	total uint32 // 票的总数
	ticketCh chan struct{} // 票的容器
	active bool // 票池是否已经被激活
}
// 注意 这里返回的是一个接口类型
func NewGoTickets(total uint32)(GoTickets,error){
	gt := myGoTickets{}
	if  !gt.init(total){
		errMsg := fmt.Sprintf("票池没有正确初始化（总数 = %d", total)
		return nil, errors.New(errMsg)
	}
	return  &gt, nil
}

func (gt *myGoTickets) init(total uint32)(bool){
	if gt.active{
		return false
	}
	if total == 0 {
		return false
	}
	ch := make(chan struct{}, total)
	n:= int(total)
	for i := 0;i < n; i++{
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}

func (gt *myGoTickets) Take(){
	<-gt.ticketCh
}
func (gt *myGoTickets) Active() bool{
	return gt.active
}
func (gt *myGoTickets) Total() uint32 {
	return gt.total
}
func (gt *myGoTickets) Return() {
	gt.ticketCh <- struct{}{}
}
func (gt *myGoTickets) Remainder()uint32{
	return uint32(len(gt.ticketCh))
}