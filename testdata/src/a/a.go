package a

func f(ch chan int) { // want "channel argument should be directed"
	// The pattern can be written in regular expression.
}

func g(ch <-chan string) {}

func h(ch chan<- float64) {}

func i(ch <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for v := range ch {
			result <- v
		}
	}()
	return result
}

func j(ch <-chan int) chan int { // want "channel result should be directed"
	result := make(chan int)
	go func() {
		for v := range ch {
			result <- v
		}
	}()
	return result
}
