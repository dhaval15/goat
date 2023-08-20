package main

import (
	"net/http"
	"goat/database"
	"goat/api"
)

func main() {
	database.NewDatabase()          // Initialize your database
	// router := api.NewRouter()      // Create router with database instance
	http.HandleFunc("/api/upload", api.UploadHandler)
	http.HandleFunc("/api/converted/{id:[0-9]+}", api.ServeConvertedFileHandler)
	fs := http.FileServer(http.Dir("./ui/build"))
	http.Handle("/", fs)
	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
