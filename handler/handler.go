package handler

import (
	"MyInternatStorage/meta"
	"MyInternatStorage/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		data, err := ioutil.ReadFile("./static/view/index.html")

		if err != nil{
			io.WriteString(w, "internal Error")
			return
		}

		io.WriteString(w, string(data))
	}else if r.Method == "POST"{
		fmt.Print("Start post !!!!")
		file, head, err := r.FormFile("file")

		if err != nil{
			fmt.Print("Failed to get data, err %s", err.Error())
			return
		}

		defer file.Close()


		// crate new filemeta for this upload file
		fileMeta := meta.FileMeta{
			Filename: head.Filename,
			Location: "/home/alan/myDir/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),


		}

		newFile, err := os.Create(fileMeta.Location)

		if err != nil{
			fmt.Print("Failed to create file, err:%s\n", err.Error())
			return

		}

		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil{
			fmt.Print("failet to save data into file %s", err.Error())
			return
		}
		//equal init the NewFile
		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		//meta.UpdateFileMeta(fileMeta)
		meta.UpdateFileMetaDb(fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}


func UploadSucHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		io.WriteString(w, "Upload Finished")
	}
}


func GetFileMetahandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	//fMeta := meta.GetFileMeta(filehash)
	fMeta, err := meta.GetFileMetaDB(filehash)

	data, err := json.Marshal(fMeta)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)


}

func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")


	if opType != "0"{
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method != "POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.Filename = newFileName
	meta.UpdateFileMeta(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)


}

func FileDeleteHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	fileSha1 := r.Form.Get("filehash")
	
	fMeta := meta.GetFileMeta(fileSha1)

	os.Remove(fMeta.Location)
	meta.RemoveFileMeta(fileSha1)
	w.WriteHeader(http.StatusOK)
}