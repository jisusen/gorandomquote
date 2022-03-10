package api

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.SetUpRoutes()
	return s
}

type Server struct {
	Router *mux.Router
}

func (s *Server) Run(ctx context.Context) error {
	srv := http.Server{
		Handler: s.Router,
		Addr:    fmt.Sprintf(":%d", 8080),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	/* Gracefull Shutdown server */
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal().Err(err).Msgf("Server Shutdown Failed")
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	return nil
}
func (s *Server) SetUpRoutes() error {
	v1 := s.Router.PathPrefix("/api/v1/").Subrouter()
	v1.HandleFunc("/quotes/random", s.handleQuote()).Methods(http.MethodGet)

	return nil
}

func (s *Server) handleQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textMap := map[string]int{
			"Foul tarnished, in search of the Elden Ring. Emboldened by the flame of ambition. Someone must extinguish thy flame. - Found at: https://www.thyquotes.com/elden-ring/":                                                                   1,
			"They will fight and they will die... in an unending curse - Found at: https://www.thyquotes.com/elden-ring/":                                                                                                                              2,
			"No one will hold me captive. A serpent never dies. - Found at: https://www.thyquotes.com/elden-ring/":                                                                                                                                     3,
			"Brandish the Elden Ring for all of us! - Found at: https://www.thyquotes.com/elden-ring/":                                                                                                                                                 4,
			"The fallen leaves tell a story. Of how a Tarnished became Elden Lord. In our home, across the fog, the Lands Between. Our seed will look back upon us, and recall. An Age of Fracture. - Found at: https://www.thyquotes.com/elden-ring/": 5,
		}
		quote := GetRandomText(textMap).(string)
		w.Write([]byte("Quotes : " + quote))
	}
}

func GetRandomText(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()
	return keys[rand.Intn(len(keys))].Interface()
}
