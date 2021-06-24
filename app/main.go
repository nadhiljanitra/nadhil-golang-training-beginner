package main

import (
	"fmt"
	"os"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/app/cmd"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "rest":
		cmd.InitRest()
	case "expirecode":
		cmd.InitExpireCode()
	default:
		fmt.Println(fmt.Errorf("Unknown argument"))
	}
}
