package crc

import (
	"testing"
)

func TestSumValue(t *testing.T) {
	if CheckSumValue([]byte{0xaa, 0xaa}) != 0xac {
		t.Fail()
	}
	if CheckSumValue([]byte{0xaa, 0xaa, 0xac}) != 0 {
		t.Fail()
	}
}
