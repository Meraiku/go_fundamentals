package main

import (
	"fmt"

	"github.com/meraiku/go_fundamentals/internal/helpers"
	"github.com/meraiku/go_fundamentals/internal/tasks"
)

func main() {

	channels := helpers.Channels()

	rCh := make([]<-chan int, len(channels))
	for i := range channels {
		rCh[i] = channels[i]
	}

	mergedCh := tasks.MergeChannels(rCh...)

	go helpers.WriteToChannels(10, channels...)

	for v := range mergedCh {
		fmt.Println(v)
	}
}
