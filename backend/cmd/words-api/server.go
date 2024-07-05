package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"

	"github.com/mrasoolmirzaei/words/backend/pkg/server"
)

func serverAction(log *logrus.Logger) cli.ActionFunc {
	return func(cc *cli.Context) error {

		serverLogger := log.WithField("context", "server")
		config := &server.Config{
			Logger:     serverLogger,
			CliContext: cc,
		}

		log.Debug("Initializing server.")
		s, err := server.NewServer(config)
		if err != nil {
			return err
		}

		g, ctx := errgroup.WithContext(context.Background())

		g.Go(func() error {
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

			select {
			case sig := <-sigChan:
				log.Infof("Received signal, exiting: %s", sig)
				return s.Stop()
			case <-ctx.Done():
				log.Infof("Received context cancel signal, exiting: %s", ctx.Err())
				return s.Stop()
			}
		})

		g.Go(func() error {
			log.Infof("Starting server on %s", cc.String("listen-http"))
			return s.Serve(cc.String("listen-http"))
		})

		return g.Wait()
	}
}
