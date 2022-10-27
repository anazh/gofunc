package date_time

import "time"

// get now unix timestamp
// second unit
func NowUnixTimeSecond() int64 {
	return time.Now().Unix()
}

// second/nano = 1e9
func NowUnixTimeNano() int64 {
	return time.Now().UnixNano()
}
