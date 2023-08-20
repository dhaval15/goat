package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/upload", UploadHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/converted/{id:[0-9]+}", ServeConvertedFileHandler).Methods(http.MethodGet)
	fs := http.FileServer(http.Dir("./ui/build"))
	router.Handle("/", fs)
	// Define other routes here
	return router
}
