package resp

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// The signature of the RESP protocol
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

// Reply is Data kind interface
type Reply interface {
	replyNode()
	Format(level uint8) string
}

// ReplyBulk Is could be any data of reply
type ReplyBulk []byte

// ReplyMultiBulk Is an array of multiple reply
type ReplyMultiBulk []Reply

// ReplyInteger is has to be integer of reply
type ReplyInteger []byte

// ReplyStatus is has to be a row of status reply
type ReplyStatus []byte

// ReplyError is has to be a row of error reply
type ReplyError []byte

func (r ReplyBulk) replyNode()      {}
func (r ReplyMultiBulk) replyNode() {}
func (r ReplyInteger) replyNode()   {}
func (r ReplyStatus) replyNode()    {}
func (r ReplyError) replyNode()     {}

func (r ReplyBulk) Format(level uint8) string {
	return strconv.Quote(*(*string)(unsafe.Pointer(&r)))
}
func (r ReplyMultiBulk) Format(level uint8) string {
	ss := make([]string, 0, len(r))
	lev := strings.Repeat("   ", int(level))
	for i, v := range r {
		text := v.Format(level + 1)
		le := lev
		if i == 0 {
			le = ""
		}
		ss = append(ss, fmt.Sprintf("%s%d) %s", le, i, text))
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
