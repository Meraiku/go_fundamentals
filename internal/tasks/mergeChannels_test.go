package tasks

import (
	"testing"

	"github.com/meraiku/go_fundamentals/internal/helpers"
	"github.com/stretchr/testify/require"
)

func TestMergeChannels(t *testing.T) {

	channels := helpers.Channels()

	rCh := make([]<-chan int, len(channels))
	for i := range channels {
		rCh[i] = channels[i]
	}

	dataPerChan := 100

	mergedCh := MergeChannels(rCh...)

	go helpers.WriteToChannels(dataPerChan, channels...)

	got := make([]int, 0, dataPerChan*len(channels))

	for v := range mergedCh {
		got = append(got, v)
	}

	require.Equal(t, dataPerChan*len(channels), len(got))
}
