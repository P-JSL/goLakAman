package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token string `Json:"Token"`
	BotPrefix string `Json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading form config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(file))

	err = json.Unmarshal(file,&config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}