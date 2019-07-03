package go_timestamp

import (
	"time"
)

var (
	// LocalTime when convert to time.Time
	LocalTime = time.Local
)

// Define a timestamp
type TimeStamp int64

// Now returns now int64
func Now() TimeStamp {
	return TimeStamp(time.Now().Unix())
}

// Add adds seconds and return sum
func (ts TimeStamp) Add(seconds int64) TimeStamp {
	return ts + TimeStamp(seconds)
}

// AddDuration adds time.Duration and return sum
func (ts TimeStamp) AddDuration(interval time.Duration) TimeStamp {
	return ts + TimeStamp(interval/time.Second)
}

// Year returns the time's year
func (ts TimeStamp) Year() int {
	return ts.AsTime().Year()
}

// Month returns the time's month
func (ts TimeStamp) Month() time.Month {
	return ts.AsTime().Month()
}

// Day returns the time's day
func (ts TimeStamp) Day() int {
	return ts.AsTime().Day()
}

// AsTime convert timestamp as time.Time in Local locale
func (ts TimeStamp) AsTime() (tm time.Time) {
	tm = time.Unix(int64(ts), 0).In(LocalTime)
	return
}

// AsTimePtr convert timestamp as *time.Time in Local locale
func (ts TimeStamp) AsTimePtr() *time.Time {
	tm := time.Unix(int64(ts), 0).In(LocalTime)
	return &tm
}

// Format formats timestamp as
func (ts TimeStamp) Format(f string) string {
	return ts.AsTime().Format(f)
}

// FormatTZ formats as datetime and zone
func (ts TimeStamp) FormatTZ() string {
	return ts.Format("20060102T150405Z")
}

// FormatDT formats as datetime
func (ts TimeStamp) FormatDT() string {
	return ts.Format("2006-01-02 15:04:05")
}

// FormatD formats as date
func (ts TimeStamp) FormatD() string {
	return ts.Format("2006-01-02")
}

// FormatT formats as time
func (ts TimeStamp) FormatT() string {
	return ts.Format("15:04:05")
}

// IsZero is zero time
func (ts TimeStamp) IsZero() bool {
	return ts.AsTime().IsZero()
}
