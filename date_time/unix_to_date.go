package date_time

import (
	"strconv"
	"time"
)

// get now second timestamp and change format to string
func NowUnixTimeString() string {
	timestamp := NowUnixTimeSecond()
	times := strconv.FormatInt(timestamp, 10)
	return times
}

// get now nano timestamp and change format to string
func NowUnixTimeStringNano() string {
	timestamp := NowUnixTimeNano()
	times := strconv.FormatInt(timestamp, 10)
	return times
}

// decode timestamp to date
func TimeDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02")
}

// decode timestamp to time
func TimeDateMinSec(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func TimeStringToUnix(timeString string) int64 {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeString, time.Local)
	out := t.Unix()
	return out
}

func DateStringToUnix(timeString string) int64 {
	t, _ := time.ParseInLocation("2006-01-02", timeString, time.Local)
	out := t.Unix()
	return out
}

func TimeYear(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006")
}

func TimeMonth(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("01")
}

func TimeDay(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("02")
}

func TimeHour(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("15")
}

func TimeMin(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("04")
}

func TimeSec(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("05")
}

func NowWeek() int {
	t := time.Now()
	return int(t.Weekday())
}
