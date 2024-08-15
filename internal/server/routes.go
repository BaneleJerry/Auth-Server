package server

import (
	"Auth-Server/internal/handler"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	db := handler.NewDBConfig(s.Database)
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", db.SignUpHandler)
		r.Post("/login", db.LoginHandler)
	})

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
