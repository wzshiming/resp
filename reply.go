package resp

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
