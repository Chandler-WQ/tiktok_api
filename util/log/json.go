package log

import (
	"github.com/bytedance/sonic"
)

func NewLogString(v interface{}) LogString {
	return LogString{
		v: v,
	}
}

type LogString struct {
	v interface{}
}

func (l LogString) String() string {
	str, _ := sonic.MarshalString(l.v)
	return str
}
