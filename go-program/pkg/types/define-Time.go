package types

import "time"

// UTC 时间格式
func AmzDate() string {
	return time.Now().UTC().Format("20060102T150405Z")
}
