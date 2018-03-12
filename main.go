package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BaritoLog/barito-flow/receiver"
	"github.com/BaritoLog/barito-flow/forwarder"
	"github.com/BaritoLog/go-boilerplate/app"
	"github.com/urfave/cli"
)

const (
	Name    = "barito-agent"
	Version = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Usage = "Provide kafka reciever or log forwarder for Barito project"
	app.Version = Version
	app.Commands = []cli.Command{
		{Name: "receiver", Usage: "Kafka Receiver", Aliases: []string{"r"}, Action: startReceiver},
		{Name: "forwarder", Usage: "Log Forwarder", Aliases: []string{"f"}, Action: startForwarder},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(fmt.Sprintf("Some error occurred: %s", err.Error()))
	}
}

func startReceiver(c *cli.Context) (err error) {
	runner := app.NewRunner(
		receiver.NewContext(),
		receiver.NewConfigurationManager(),
	)

	err = runner.Run()
	return
}

func startForwarder(c *cli.Context) (err error) {
	runner := app.NewRunner(
		forwarder.NewContext(),
		forwarder.NewConfigurationManager(),
	)

	err = runner.Run()

	return
}
