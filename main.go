package main

import (
	"fmt"
	"os"

	"github.com/clbx/plantcam/reolink"
)

func main() {
	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		fmt.Printf("ADDRESS environment variable not set\n")
	}

	user, exists := os.LookupEnv("U")
	if !exists {
		fmt.Printf("USERNAME environment variable not set\n")
	}

	pass, exists := os.LookupEnv("PASS")
	if !exists {
		fmt.Printf("PASSWORD environment variable not set\n")
	}

	fmt.Printf("%s %s %s\n",address,user,pass)

	reolink.Login(address,user,pass)
}
