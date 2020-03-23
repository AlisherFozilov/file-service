package app

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

const multipartMaxBytes = 100 * 1024 * 1024

/*
func NewFileHandler(svc *files.FileService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
*/

func (s *Server) handleMultipartUpload(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		return
	}

	err := request.ParseMultipartForm(multipartMaxBytes)
	if err != nil {
		log.Print(err)
		http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fileHeaders := request.MultipartForm.File["file"]

	fileURLs := make([]FileURL, 0, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
		name, err := s.saveFile(fileHeader)
		if err != nil {
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Print(err)
			return
		}

		fileURLs = append(fileURLs, FileURL{
			Id:   name[:len(name)-len(filepath.Ext(name))],
			Path: "http://" + request.Host + "/" + s.storagePath + "/" + name,
		})
	}

	urlsJSON, err := json.Marshal(fileURLs)
	if err != nil {
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	_, err = responseWriter.Write(urlsJSON)
	if err != nil {
		log.Print(err)
		return
	}

	return
}

/*
func (s *Server) handleUploadPage(responseWriter http.ResponseWriter, request *http.Request) {

}
*/

func (s *Server) saveFile(fileHeader *multipart.FileHeader) (name string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer func() {
		errdefer := file.Close()
		if errdefer != nil {
			log.Println(errdefer)
		}
	}()

	contentType := fileHeader.Header.Get("Content-Type")
	name, err = s.fileSvc.Save(file, contentType)
	if err != nil {
		return
	}

	return //nil
}
