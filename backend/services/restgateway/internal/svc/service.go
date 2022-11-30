package svc

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Service struct {
	routes http.Handler
}

func (s *Service) Init() {
	router := mux.NewRouter()
	s.routes = router

	buildRoutes(router)
}

func (s *Service) Start() {
	fmt.Println("Service started at port 8080")
	http.ListenAndServe(":8080", s.routes)
}

func buildRoutes(router *mux.Router) {
	path, err := filepath.Abs("./swaggerui")
	if err != nil {
		panic(err)
	}

	router.Use(mux.CORSMethodMiddleware(router))

	swaggerUIHandler := http.StripPrefix("/swaggerui", http.FileServer(http.Dir(path)))

	router.PathPrefix("/swaggerui").Methods(http.MethodGet).Handler(swaggerUIHandler)
}
