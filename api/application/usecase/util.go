package usecase

import (
	"time"

	"github.com/myapp/noname/api/env"
)

func init() {
	timeZoneSetJST()
}

func timeZoneSetJST() {
	if env.IsStaging() && env.IsProduction() {
		time.Local = time.FixedZone("Local", -8*60*60)
	}
}
