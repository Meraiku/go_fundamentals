package helpers

func Channels() []chan int {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)
	return []chan int{ch1, ch2, ch3, ch4, ch5}
}

func WriteToChannels(ammount int, channels ...chan int) {

	for _, ch := range channels {
		go func() {
			for i := 0; i < ammount; i++ {
				ch <- i + 1
			}
			close(ch)
		}()
	}

}
