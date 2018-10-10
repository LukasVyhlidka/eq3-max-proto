package main

import (
	"fmt"

	"github.com/LukasVyhlidka/eq3-max-proto/model"
)

func main() {
	fmt.Println("Hello")

	msg := model.MaxDevice{
		RfAddress: 1,
	}

	fmt.Printf("Hello - %d\n", msg.RfAddress)
	fmt.Print(msg)
}
