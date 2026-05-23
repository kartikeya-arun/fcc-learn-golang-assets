package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int = 0
	var costPerSMS float32 = 0.0
	var hasPermission bool = false
	var username string = ""
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
