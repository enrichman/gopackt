package gopackt

import (
	"fmt"
	"time"
)

func FancyLoad(done *bool, doneChan chan bool) {
	for !*done {
		fmt.Print("\r[\\]")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\r[|]")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\r[/]")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\r[-]")
		time.Sleep(50 * time.Millisecond)
	}
	doneChan <- true
}
