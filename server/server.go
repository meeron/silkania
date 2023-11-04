package server

import (
	"fmt"
	"time"
)

var StartTime time.Time

func Uptime() string {
	return fmt.Sprintf("%v", time.Since(StartTime))
}
