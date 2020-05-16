package main

import (
	"fmt"
	
	"github.com/roikramer120/autogo/pkg/slacknotify"
)

func main() {
	fmt.Println("test")
	sn, err := slacknotify.Init(nil)
	sn.SendString("Hello channel")
}