package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yuhang2/mockdata/config"
	"github.com/yuhang2/mockdata/operation"
)

func main() {
	initServer()

	opt := flag.String("opt", "", "location - get location; booking - generate booking data")
	line := flag.Int("line", 0, "number of location")
	flag.Parse()

	switch *opt {
	case "location":
		if *line <= 0 {
			break
		}
		operation.Location(*line)
	case "booking":
		operation.Booking()
	}
	fmt.Println("It's over!")
}

func initServer() {
	err := config.NewConfig()
	if err != nil {
		log.Fatal("initServer error")
		panic(err)
	}
}
