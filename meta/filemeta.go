package meta

import (
	"MyInternatStorage/db"
	"fmt"
)



type FileMeta struct{
	FileSha1 string
	Filename string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta


func init(){
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmeta FileMeta){
	fileMetas[fmeta.FileSha1] = fmeta
}


func GetFileMeta(fileSha1 string)FileMeta{
	return fileMetas[fileSha1]
}


func RemoveFileMeta(fileSha1 string){
	delete(fileMetas, fileSha1)
}


func UpdateFileMetaDb(fmeta FileMeta)bool{
	fmt.Println(fmeta.FileSha1)
	fmt.Println(fmeta.FileSize)

	fmt.Println(fmeta.Filename)

	fmt.Println(fmeta.Location)


	return db.OnFileUploadFinished(fmeta.FileSha1, fmeta.Filename, fmeta.FileSize, fmeta.Location)
}

func GetFileMetaDB(filehash string) (FileMeta, error) {

		tfile, err := db.GetFileMeta(filehash)
		if err != nil{
			return FileMeta{}, err
		}

		fmeta:=FileMeta{
			FileSha1: tfile.FileHash,
			Filename: tfile.FileName.String,
			FileSize: tfile.FileSize.Int64,
			Location: tfile.FileAddr.String,
		}

		return fmeta, nil
}