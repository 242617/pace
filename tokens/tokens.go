package tokens

import (
	"log"
	"time"
)

const (
	TokenTypeSMSRequest = iota
)

func NewToken(mode int, claims interface{}) string {

	var duration time.Duration
	switch mode {
	case TokenTypeSMSRequest:
		duration = time.Minute
	default:
		duration = time.Minute
	}

	log.Println(duration)

	return ""
}
