package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func UploadFile(w http.ResponseWriter, r * http.Request) {
	//  Ensure our file does not exceed 5MB
	r.Body = http.MaxBytesReader(w, r.Body, 5 * 1024 * 1024)
  
	file, handler, err := r.FormFile("image")
  
	// Capture any errors that may arise
	if err != nil {
	  fmt.Fprintf(w, "Error getting the file")
	  fmt.Println(err)
	  return
	}
  
	defer file.Close()
  
	fmt.Printf("Uploaded file name: %+v\n", handler.Filename)
	fmt.Printf("Uploaded file size %+v\n", handler.Size)
	fmt.Printf("File mime type %+v\n", handler.Header)
  
	// Get the file content type and access the file extension
	fileType := strings.Split(handler.Header.Get("Content-Type"), "/")[1]
  
	// Create the temporary file name
	fileName := fmt.Sprintf("upload-*.%s", fileType)
	// Create a temporary file with a dir folder
	tempFile, err := ioutil.TempFile("temp-files", fileName)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	defer tempFile.Close()
  
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
	  fmt.Println(err)
	}
  
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully uploaded file")
}