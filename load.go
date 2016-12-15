package gopackt

import (
	"fmt"
	"time"
)

func FancyLoad(quit chan bool) {
	wheel := []string{"\\", "|", "/", "-"}
	var i int
	for {
		select {
		case <-quit:
			close(quit)
			return
		default:
			fmt.Print("\r[" + wheel[i] + "]")
			i = (i + 1) % len(wheel)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Print("\033[2K\r")
	}
}
