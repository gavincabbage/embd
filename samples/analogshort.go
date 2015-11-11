// +build ignore

package main

import (
	"flag"
	"fmt"

	"github.com/gavincabbage/embd"

	_ "github.com/gavincabbage/embd/host/all"
)

func main() {
	flag.Parse()

	embd.InitGPIO()
	defer embd.CloseGPIO()

	val, _ := embd.AnalogRead(0)
	fmt.Printf("Reading: %v\n", val)
}
