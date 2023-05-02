package helps

import (
	"time"
)

const ONE_MINUTE = 60
const THREE_MINUTE = 181

func CheckTokenExpired(expiredTime int64) bool {
	now := time.Now().Unix()
	return now-expiredTime > 0
}

func GetCurrentMicroSecond() int64 {
	return time.Now().UnixNano() / 1000
}

func CheckTokenExpiredSecond(timeSign int64) bool {
	now := time.Now().Unix()
	return (now-THREE_MINUTE)-timeSign > 0
}
