package gopackt

import (
	"fmt"
	"time"
)

func FancyLoad(quit chan bool) {
	defer close(quit)

	wheel := []string{"\\", "|", "/", "-"}
	var i int
	for {
		select {
		case <-quit:
			return
		default:
			fmt.Print("\r[" + wheel[i] + "]")
			i = (i + 1) % len(wheel)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Print("\033[2K\r")
	}
}
