package tokens

import (
	"fmt"
	"time"
)

const (
	TokenModeSMSRequest = iota
)

func NewToken(mode int, claims interface{}) string {

	var duration time.Duration
	switch mode {
	case TokenModeSMSRequest:
		duration = time.Minute
	default:
		duration = time.Minute
	}

	fmt.Println(duration)

	return ""
}
