package helpers

import "time"

func GetCurrentTime() time.Time {
	return time.Now().In(time.FixedZone("GMT+5", 7*60*60))
}
