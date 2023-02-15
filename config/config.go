package config

import (
	"encoding/json"
	"log"
	"io/ioutil"
	)

type Config struct {
	ServerAddress      string `json:"serverAddress"`
	UDPPort			   int	  `json:"UDPPort"`
	TCPPort           int    `json:"TCPPort"`

}

var (
	G_Config *Config
)

func InitConfig(fileName string) (err error) {
	var (
		content []byte
		conf    Config
	)
	if content, err = ioutil.ReadFile(fileName); err != nil {
		log.Println("failed to read configuration file: ", fileName, ", ", err.Error())
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		log.Println("failed to  Unmarshal configuration file:", err.Error())
		return
	}

	G_Config = &conf

	log.Println(G_Config)
	return
}