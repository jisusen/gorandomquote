package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}

	return s
}

type Server struct {
	Router *mux.Router
}

func (s *Server) runServer() error {
	srv := http.Server{
		Handler: s.Router,
		Addr:    fmt.Sprintf(":%d", ":8080"),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()

	/* Gracefull Shutdown server */
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal().Err(err).Msgf("Server Shutdown Failed")
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	return nil
}
func (s *Server) routes() error {
	v1 := s.Router.PathPrefix("/api/v1/").Subrouter()

	v1.HandleFunc("/quotes/random", s.handleQuote()).Methods(http.MethodGet)

	return nil
}

func (s *Server) handleQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oce"))
	}
}
