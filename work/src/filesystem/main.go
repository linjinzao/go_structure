package main

import(
	"net/http"
	handle "filesystem/handle"
	"log"
)

func main(){

	log.Println("server start!")
	http.HandleFunc("/file/upload",handle.UploadHandler)
	http.HandleFunc("/file/upload/suc",handle.UploadSucHandler)
	http.HandleFunc("/file/meta",handle.GetFileMeta)
	http.HandleFunc("/file/download",handle.DownloadHandler)
	http.HandleFunc("/file/update",handle.FileUpdateMetaHandler)
	http.HandleFunc("/file/delete",handle.FileDeleteHandler)
	http.ListenAndServe("127.0.0.1:8088",nil)
}