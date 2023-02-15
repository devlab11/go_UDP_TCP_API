package main

import (
	"fmt"
	"log"	
	"time"
	"flag"	
	"go_testtask/server"
	"go_testtask/config"
)

var (
	confFilePath string
)

func initArgs() {
	flag.StringVar(&confFilePath, "config", "config/config.json", "Configfile path")
	flag.Parse()
}	

func main() {
	var (
		err    error		
	)

	initArgs()
    fmt.Println("init Args!")

	if err = config.InitConfig(confFilePath); err != nil {
		log.Fatal ("failed to load configuration: ", err.Error())
		goto ERR
	}
	log.Println("Initial configuration  success")
	
	server.InitUDPServer()
	log.Println("Initial UDP server success")

	go server.G_UDPServer.CreateListener()

	for {
		time.Sleep(1*time.Second)
	}
ERR:
log.Fatalln(err.Error())
}