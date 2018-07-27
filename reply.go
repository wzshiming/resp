package resp

import (
	"fmt"
	"strings"
	"unsafe"
)

const (
	Status    = '+'
	Error     = '-'
	Integer   = ':'
	Bulk      = '$'
	MultiBulk = '*'
)

var (
	crlf   = []byte{'\r', '\n'}
	nilVal = []byte{'-', '1'}
)

type Reply interface {
	replyNode()
	Format(level uint8) string
}

type ReplyBulk []byte
type ReplyMultiBulk []Reply
type ReplyInteger []byte
type ReplyStatus []byte
type ReplyError []byte

func (r ReplyBulk) replyNode()      {}
func (r ReplyMultiBulk) replyNode() {}
func (r ReplyInteger) replyNode()   {}
func (r ReplyStatus) replyNode()    {}
func (r ReplyError) replyNode()     {}

func (r ReplyBulk) Format(level uint8) string {
	return *(*string)(unsafe.Pointer(&r))
}
func (r ReplyMultiBulk) Format(level uint8) string {
	ss := make([]string, 0, len(r))
	lev := strings.Repeat(" ", int(level))
	for i, v := range r {
		text := v.Format(level + 1)
		ss = append(ss, fmt.Sprintf("%s%d) %s", lev, i, text))
	}
	out := strings.Join(ss, "\n")
	return out
}
func (r ReplyInteger) Format(level uint8) string {
	text := append([]byte("(Integer) "), r...)
	return *(*string)(unsafe.Pointer(&text))
}
func (r ReplyStatus) Format(level uint8) string {
	text := append([]byte("(Status) "), r...)
	return *(*string)(unsafe.Pointer(&text))
}
func (r ReplyError) Format(level uint8) string {
	text := append([]byte("(Error) "), r...)
	return *(*string)(unsafe.Pointer(&text))
}
