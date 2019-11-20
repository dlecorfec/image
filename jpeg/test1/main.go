package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("i=%d xOff=%d yOff=%d\n", i, (i&1)*8, (i&2)*4)
	}
}
