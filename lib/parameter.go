package lib

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

type ParamSet struct {
	Caller Caller
	TimeoutNS time.Duration
	LPS uint32
	DurationNS time.Duration
	ResultCh chan *CallResult
}
func (pset *ParamSet) Check() error{
	var errMsgs []string
	if pset.Caller == nil{
		errMsgs = append(errMsgs,"Invalid caller")
	}
	if pset.TimeoutNS == 0{
		errMsgs = append(errMsgs,"Invalid timeoutNS")
	}
	if pset.LPS == 0{
		errMsgs =  append(errMsgs,"Invalid lps")
	}
	if pset.DurationNS == 0{
		errMsgs = append(errMsgs, "Invalid durationNS")
	}
	if pset.ResultCh == nil{
		errMsgs = append(errMsgs, "Invalid result channel")
	}
	var buf bytes.Buffer
	buf.WriteString("Checking the parameters...")
	if errMsgs != nil{
		errMsgs := strings.Join(errMsgs," ")
		buf.WriteString(fmt.Sprintf("NOT PASSED! (%s)",errMsgs))
		log.Println(buf.String())
		return errors.New(errMsgs)
	}
	buf.WriteString(
		fmt.Sprintf("Passed. (timeoutNS = %s, lps = %d, durationNS = %s)",
			pset.TimeoutNS,pset.LPS,pset.DurationNS))
	log.Println(buf.String())
	return nil
}