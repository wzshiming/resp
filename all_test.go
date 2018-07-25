package resp

import (
	"bytes"
	"testing"
)

var testdata = []string{
	"+ok\r\n",
	"-error \r\n",
	":113213\r\n",
	"$10\r\nhello 1234\r\n",
	":-100\r\n",
	"*-1\r\n",
	"$-1\r\n",
	"*2\r\n*-1\r\n:666\r\n",
	"*4\r\n$5\r\nhello\r\n$5\r\nworld\r\n*2\r\n:100\r\n$5\r\ntimes\r\n*2\r\n+OK\r\n-Error\r\n",
}

func TestBasics(t *testing.T) {
	for _, in := range testdata {
		bufin := bytes.NewBufferString(in)
		data, err := NewDecoder(bufin).Decode()
		if err != nil {
			t.Fatal(err, in)
			return
		}

		bufout := bytes.NewBuffer(nil)
		err = NewEncoder(bufout).Encode(data)
		if err != nil {
			t.Fatal(err, in)
			return
		}
		if in != bufout.String() {
			t.Log(in)
			t.Log(bufout.String())
			t.Fatal("error")
			return
		}
	}
}
