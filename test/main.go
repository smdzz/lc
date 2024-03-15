package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	fmt.Println("Start...")
	for i := 0; i < 10000000; i++ {
		x := rand.Intn(100)
		y := rand.Intn(100)
		fmt.Println(x, y)
		_ = x + y
	}
	fmt.Println("Done!")

}
