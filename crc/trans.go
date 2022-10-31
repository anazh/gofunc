package crc

import (
	"bytes"
	"encoding/binary"
)

func Int16To2Byte(num interface{}) []byte {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, num)
	return buffer.Bytes()
}
