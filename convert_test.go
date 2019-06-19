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
		d, err := ConvertTo(v)
		if err != nil {
			t.Fatal(err)
			return
		}
		err = NewEncoder(buf).Encode(d)
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

func TestConvertAll(t *testing.T) {
	type tt struct {
		S  string
		I  int
		M  map[string]int
		Sl []uint
		Ar [2]int
		PS *string
	}
	var ps = "ps"
	b := tt{"s", 100, map[string]int{"aa": 111}, []uint{1, 2, 3, 4}, [2]int{-1, 2}, &ps}
	tmp, err := ConvertTo(b)
	if err != nil {
		t.Fatal(err)
		return
	}
	b0 := tt{}
	err = ConvertFrom(tmp, &b0)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(b0)
}
