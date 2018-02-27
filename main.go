package main

import (
	"fmt"
	"log"
	"logstore/receiver"
	"os"

	"github.com/urfave/cli"
)

const (
	Name    = "barito"
	Version = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Usage = "Barito Raftman Cli"
	app.Version = Version
	app.Commands = []cli.Command{
		{Name: "receiver", Usage: "Start Receiver", Aliases: []string{"r"}, Action: startReceiver},
		{Name: "forwarder", Usage: "Start Forwarder", Aliases: []string{"f"}, Action: startForwarder},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(fmt.Sprintf("Some error occurred: %s", err.Error()))
	}
}

func startReceiver(c *cli.Context) (err error) {
	fmt.Println("Receiver")

	addr := ":8080"

	lsReceiver := receiver.NewLogstoreReceiver(addr)
	err = lsReceiver.Start()

	return
}

func startForwarder(c *cli.Context) (err error) {
	fmt.Println("Forwarder - Under Construction")

	return
}