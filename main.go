package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

var nSeconds = flag.Int64("nSeconds", 20, "How many seconds will players play")

func main() {
	flag.Parse()
	t := runtime.GOMAXPROCS(16)
	//nSeconds := 20
	fmt.Printf("%v threads are used\n", t)

	ball := make(chan struct{})
	done := make(chan struct{})
	n := 0
	/*ticker := time.NewTicker(1 * time.Second)
	<-ticker.C // Получение из канала ticker
	ticker.Stop() // Остановка go-подпрограммы ticker*/
	i := 2

	for i > 0 {
		var curi = i
		go func(int) {
			fmt.Printf("%v-player starts\n", curi)
			for {
				select {
				case _, opened := <-ball:
					if opened {
						n++
						ball <- struct{}{}
					} else {
						fmt.Printf("%v-player is done\n", curi)
						return
					}
				case <-done:
					fmt.Printf("%v-player is done\n", curi)
					return
					//default:
					//	fmt.Printf("%v-player is done\n", curi)
					//	return
				}
			}
		}(curi)
		i--
	}
	ball <- struct{}{}
	time.Sleep(time.Duration(*nSeconds) * time.Second)
	<-ball
	close(ball)
	close(done)
	fmt.Printf("%v givings per %v seconds was done\n", n, nSeconds)
	fmt.Scanln()
}
