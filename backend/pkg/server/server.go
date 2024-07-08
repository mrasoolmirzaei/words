package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mrasoolmirzaei/words/backend/pkg/api"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Server implements the main server over http.
type Server struct {
	api      api.API
	log      logrus.FieldLogger
	router   *mux.Router
	stopChan chan struct{}
}

// Config is used to initialize a new server.
type Config struct {
	Api        *api.API
	CliContext *cli.Context
	Logger     logrus.FieldLogger
}

func NewServer(c *Config) (*Server, error) {
	if c.Logger == nil {
		return nil, errors.New("logger must be specified and cannot be nil")
	}

	s := &Server{
		api:      *c.Api,
		router:   mux.NewRouter(),
		log:      c.Logger,
		stopChan: make(chan struct{}),
	}

	s.routes()
	return s, nil
}

func (s *Server) routes() {
	s.router.HandleFunc("/word", s.addWord()).Methods("POST")
	s.router.HandleFunc("/synonym/{word}", s.addSynonym()).Methods("POST")
	s.router.HandleFunc("/synonyms/{word}", s.getSynonyms()).Methods("GET")
}

func (s *Server) Serve(listen string) error {

	hs := http.Server{
		Addr:    listen,
		Handler: s.router,
	}

	go func() {
		<-s.stopChan
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		s.log.Info("Shutting down HTTP server.")
		if err := hs.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
			s.log.WithError(err).Error("failed to shutdown HTTP server")
		}
	}()

	if err := hs.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	select {
	case <-s.stopChan:
		// Already closed. Don't close again.
	default:
		// Safe to close here. We're the only closer.
		close(s.stopChan)
	}

	return nil
}
