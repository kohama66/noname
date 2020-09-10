package usecase

import (
	"time"
)

func init() {
	timeZoneSetJST()
}

func timeZoneSetJST() {
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
}
