package crc

//
func CheckSumValue(data []byte) byte {
	var v byte
	for _, d := range data {
		v = d + v
	}
	return ^(v - 1)
}
