package crc

//add sum value
func SumValue(data []byte) byte {
	v := 0
	for _, d := range data {
		v = int(d) + v
	}
	return byte(int8(v))
}
