package tasks

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetTime(t *testing.T) {

	timeNow := time.Now()

	now = func() time.Time {
		return timeNow
	}

	require.Equal(t, GetTime(), fmt.Sprintf("Current time: %s", timeNow))
}
