package app

import (
	"errors"
	"github.com/AlisherFozilov/file-service/pkg/services/files"
	"net/http"
)

type Server struct {
	router        http.Handler
	fileSvc       *files.FileService
	storagePath   string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(router http.Handler, fileSvc *files.FileService, storagePath string) *Server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if fileSvc == nil {
		panic(errors.New("fileSvc can't be nil"))
	}
	if storagePath == "" {
		panic(errors.New("storagePath can't be nil"))
	}

	return &Server{fileSvc: fileSvc, storagePath: storagePath, router: router}
}
