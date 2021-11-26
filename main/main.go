package main

import (
	"botimgs/bot"
	"botimgs/config"
	"fmt"
)

// Variables used for command line parameters

const Token string = "OTEzNjAzMTMzMDg5OTE5MDM2.YaA5OA._U9e5YChwPt9QYNXVXdY7eJcRWI"
var BotId string

func main() {

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<- make(chan struct{})
	return
}