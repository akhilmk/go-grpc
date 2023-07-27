package main

import (
	"log"

	notifclient "github.com/akhilmk/go-grpc/pushnotification/client"
	notifserver "github.com/akhilmk/go-grpc/pushnotification/server"
)

func main() {
	log.Println("main start")
	notifserver.RunServer()
	notifclient.RunClient()
	log.Println("main end")
}
