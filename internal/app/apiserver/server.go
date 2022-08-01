package apiserver

import (
	"aws/internal/app/store"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(handlers.CORS(
		handlers.AllowedMethods([]string{http.MethodGet}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
	))

	s.router.HandleFunc("/users", s.handleAcademic()).Methods("GET")
	s.router.HandleFunc("/university", s.handleUniversity()).Methods("GET")
	s.router.HandleFunc("/school{id:[0-9]+}", s.handleSchool()).Methods("GET")
	s.router.HandleFunc("/program{id:[0-9]+}", s.handleProgram()).Methods("GET")
}

func (s *server) handleAcademic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *server) handleUniversity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.store.Academic().GetUniversity()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handleSchool() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		school, err := s.store.Academic().GetSchool(vars["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		s.respond(w, r, http.StatusOK, school)
	}
}

func (s *server) handleProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		program, err := s.store.Academic().GetProgram(vars["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		s.respond(w, r, http.StatusOK, program)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
