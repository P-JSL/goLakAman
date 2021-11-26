package main

import (
	"botimgs/bot"
	"botimgs/config"
	"fmt"
)

// Variables used for command line parameters

var BotId string

func main() {

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})
	return
}
