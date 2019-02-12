package strcase

import (
	"reflect"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

func ToCamel(str string) string {
	if str == `` {
		return ``
	}
	buf := make([]byte, len(str))
	ptr := unsafe.Pointer(*(**byte)(unsafe.Pointer(&buf)))
	offset, up := 0, true
	hdr := new(reflect.SliceHeader)

	for _, r := range str {
		if r == '_' || r == '-' || unicode.IsSpace(r) {
			up = true
			continue
		}
		if up && unicode.IsLetter(r) {
			up, r = false, unicode.ToUpper(r)
		}
		if r >= utf8.RuneSelf {
			//offset += utf8.EncodeRune(buf[offset:], r)
			hdr.Data = uintptr(ptr) + uintptr(offset)
			hdr.Len = len(buf) - offset
			hdr.Cap = cap(buf) - offset
			offset += utf8.EncodeRune(*(*[]byte)(unsafe.Pointer(hdr)), r)
		} else {
			//buf[offset], offset = byte(r), offset+1
			*(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(offset))) = byte(r)
			offset++
		}
	}
	//return string(buf[:offset])
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: (uintptr)(ptr),
		Len: offset,
	}))
}
