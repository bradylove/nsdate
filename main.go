package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var duration time.Duration
	if len(os.Args) >= 2 {
		var err error
		duration, err = time.ParseDuration(os.Args[1])
		if err != nil {
			fmt.Println("failed to parse duration")
			os.Exit(1)
		}
	}

	fmt.Print(time.Now().Add(duration).UnixNano())
}
