package api

import (
	"encoding/json"
	"fmt"
	"goat/database"
	"goat/utils"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Limit the maximum uploaded file size

	title := r.FormValue("title")
	// tags := r.Form["tags"]
	converterType := r.FormValue("converter_type")

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Initialize your file converter registry here
	fileConverterRegistry := utils.FileConverterRegistry()
	fileConverter, ok := fileConverterRegistry[converterType]
	if !ok {
		http.Error(w, "Unsupported converter type", http.StatusBadRequest)
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		return
	}

	token := utils.GenerateToken() // Implement a function to generate tokens
	hashedToken := utils.HashString(token) // Implement hashString function
	document := &database.Document{
		Title:        title,
		// Tags:         tags,
		HashedToken:  hashedToken,
		ConverterType: converterType,
	}

	// Save the original file to a directory with the document ID as the filename
	err = utils.SaveFileToDirectory(fileBytes,"original_files", string(document.ID))
	if err != nil {
		http.Error(w, "Error saving original file", http.StatusInternalServerError)
		return
	}

	// Convert the file using the file converter
	convertedFileBytes, err := fileConverter.Convert(fileBytes)
	if err != nil {
		http.Error(w, "Error converting file", http.StatusInternalServerError)
		return
	}

	// Save the converted file to a directory with the document ID as the filename
	err = utils.SaveFileToDirectory(convertedFileBytes,"converted_files", string(document.ID))
	if err != nil {
		http.Error(w, "Error saving converted file", http.StatusInternalServerError)
		return
	}

	db := database.NewDatabase()
	err = db.SaveDocument(document)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, "Error saving document to database", http.StatusInternalServerError)
		return
	}

	// Return the token and ID as the response
	responseData := map[string]interface{}{
		"id":    document.ID,
		"token": token,
	}
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func ServeConvertedFileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	token := r.Header.Get("Authorization")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Initialize your database instance and retrieve token by ID
	db := database.NewDatabase()
	document, err := db.FindDocumentByID(uint(id), token)
	_ = document
	if err != nil {
		http.Error(w, "Invalid ID or document not found", http.StatusUnauthorized)
		return
	}

	// ... Continue with serving the converted file
}

