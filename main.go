package main

import (
	"MyInternatStorage/handler"
	"fmt"
	"net/http"
)

func main(){
	//http.HandleFunc("/file/upload", handler.UploadHandler)
	//http.HandleFunc("/file/upload/suc", handler.Upload)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetahandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)


	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Print("Failed to start server")
	}
}
