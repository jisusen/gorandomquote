package main

import (
	"context"
	"github.com/jisusen/gorandomquote/backend/api"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := api.NewServer()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		oscall := <-ch
		log.Info().Msgf("System call: %v", oscall)
		cancel()
	}()

	err := s.Run(ctx)
	if err != nil {
		log.Warn().Msgf("Error: %v", err)
	}
}