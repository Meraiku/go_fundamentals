package tasks

import (
	"fmt"
	"time"
)

var now = time.Now

func GetTime() string {
	return fmt.Sprintf("Current time: %s", now())
}
