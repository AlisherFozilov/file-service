package main

import (
	"flag"
	"github.com/AlisherFozilov/file-service/cmd/file-server/app"
	"github.com/AlisherFozilov/file-service/pkg/services/files"
	"net"

	"log"
	"net/http"
	"os"
)

var storagePathPtr = flag.String("storage", "serverdata", "folder for storing files")
var host = flag.String("host", "0.0.0.0", "Server host")
var port = flag.String("port", "9999", "Server port")

func main() {
	flag.Parse()
	start()
}

func start() {
	createStorageFolder()
	mux := http.NewServeMux()
	fileSvc := files.NewFilesSvc(files.StoragePathType(*storagePathPtr))
	server := app.NewServer(
		mux,
		fileSvc,
		*storagePathPtr,
	)

	server.InitRoutes()
	log.Fatal(http.ListenAndServe(net.JoinHostPort(*host, *port), server))
}

func createStorageFolder() {
	err := os.Mkdir(*storagePathPtr, 0666)
	if err != nil {
		if !os.IsExist(err) {
			log.Fatalf("can't create directory: %s", err)
		}
	}
}
