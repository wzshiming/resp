package resp

import (
	"bytes"
	"testing"
)

var convertTestData = map[string]interface{}{
	"*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n":   []string{"hello", "world"},
	"*3\r\n$3\r\nget\r\n$3\r\nkey\r\n:1\r\n": []interface{}{"get", []byte("key"), 1},
}

func TestConvert(t *testing.T) {
	for k, v := range convertTestData {
		buf := bytes.NewBuffer(nil)
		err := NewEncoder(buf).Encode(Convert(v))
		if err != nil {
			t.Fatal(err)
			return
		}
		data := buf.String()
		if k != data {
			t.Fatal(k, data)
			return
		}
	}
}
