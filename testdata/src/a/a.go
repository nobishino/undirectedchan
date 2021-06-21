package a

func f(ch chan int) { // want "channel argument should be directed"
	// The pattern can be written in regular expression.
}

func g(ch <-chan string) {}

func h(ch chan<- float64) {}
