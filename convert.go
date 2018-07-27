package resp

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// Convert returns convert the base type to Reply
func Convert(i interface{}) Reply {
	val := reflect.ValueOf(i)
	return convert(val)
}

func convert(val reflect.Value) Reply {
	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return ReplyInteger{'1'}
		}
		return ReplyInteger{'0'}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return ReplyInteger(strconv.AppendInt(nil, val.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return ReplyInteger(strconv.AppendUint(nil, val.Uint(), 10))
	case reflect.Float32:
		return ReplyBulk(strconv.AppendFloat(nil, val.Float(), 'f', -1, 32))
	case reflect.Float64:
		return ReplyBulk(strconv.AppendFloat(nil, val.Float(), 'f', -1, 64))
	case reflect.Complex64, reflect.Complex128:
		com := fmt.Sprint(val.Complex())
		return ReplyBulk(*(*[]byte)(unsafe.Pointer(&com)))
	case reflect.Interface, reflect.Ptr:
		return convert(val.Elem())
	case reflect.Map:
		num := val.Len()
		items := make(ReplyMultiBulk, 0, num*val.Len())

		for _, k := range val.MapKeys() {
			items = append(items, convert(k))
			items = append(items, convert(val.MapIndex(k)))
		}
		return items
	case reflect.Struct:
		num := val.NumField()
		typ := val.Type()
		items := make(ReplyMultiBulk, 0, num*2)
		for i := 0; i != num; i++ {
			name := typ.Field(i).Name
			items = append(items, ReplyBulk(*(*[]byte)(unsafe.Pointer(&name))))
			items = append(items, convert(val.Field(i)))
		}
		return items
	case reflect.Slice:
		switch val.Type().Elem().Kind() {
		case reflect.Uint8:
			return ReplyBulk(val.Bytes())
		}
		num := val.Len()
		items := make(ReplyMultiBulk, 0, num)
		for i := 0; i != num; i++ {
			items = append(items, convert(val.Index(i)))
		}
		return items
	case reflect.Array:
		num := val.Len()
		items := make(ReplyMultiBulk, 0, num)
		for i := 0; i != num; i++ {
			items = append(items, convert(val.Index(i)))
		}
		return items
	case reflect.String:
		com := val.String()
		return ReplyBulk(*(*[]byte)(unsafe.Pointer(&com)))
	case reflect.UnsafePointer:
		return ReplyInteger(strconv.AppendUint(nil, uint64(val.Pointer()), 10))
	}
	return ReplyError(fmt.Sprintf("Error Data types are not supported '%s'", val.Kind().String()))
}
