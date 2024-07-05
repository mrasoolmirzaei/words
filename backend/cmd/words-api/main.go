package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	cmd := cli.NewApp()
	cmd.Name = "words-backend"

	log := logrus.New()
	log.Out = os.Stdout

	log.Info("Starting words-backend service...")

	cmd.Commands = []*cli.Command{
		{
			Name:   "server",
			Flags:  getServerFlags(),
			Action: serverAction(log),
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.WithFields(logrus.Fields{
			"context": "main",
			"version": cmd.Version,
		}).Fatalln(err)
	}
}

func getServerFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "listen-http",
			Usage:   "listen HTTP `IP:PORT`",
			EnvVars: []string{"LISTEN_HTTP"},
			Value:   ":8080",
		},
	}
	return flags
}
