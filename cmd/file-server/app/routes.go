package app

import "net/http"

func (s *Server) InitRoutes() {
	mux := s.router.(*http.ServeMux)

	mux.HandleFunc("/api/files", s.handleMultipartUpload)
	mux.Handle("/" + s.storagePath + "/",
		http.StripPrefix("/" + s.storagePath + "/", http.FileServer(http.Dir(s.storagePath))))
}
