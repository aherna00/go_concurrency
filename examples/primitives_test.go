package examples

import (
	"fmt"
	"testing"
	"time"
)

func doSomething() error {
	return fmt.Errorf("uh oh")
}

func TestDoSomething(t *testing.T) {

	//demonstrate go routine

	//synchronize these go routines

	doneChan := make(chan bool)

	go func() {
		err := doSomething()
		fmt.Println(err)

		doneChan <- true

	}()

	err := doSomething()

	fmt.Println(err)

	<-doneChan

}

func TestChannels(t *testing.T) {

	//create channel
	//demonstrate blocking capability

}

func TestSelect(t *testing.T) {

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			ch1 <- "test"
		}

	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			ch2 <- "test2"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch3 <- "test3"
		}
	}()

	for {
		select {
		case <-ch1:
			fmt.Println("case 1 reached")

		case <-ch2:
			fmt.Println("case 2 reached")

		case <-ch3:
			fmt.Println("case 3 reached")
		}

	}

}
