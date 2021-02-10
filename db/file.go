package db

import (
	mydb "MyInternatStorage/db/mysql"
	"database/sql"
	"fmt"
)


func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {



	stmt, err:= mydb.DBConn().Prepare(
			"insert ignore into tbl_file (`file_sha1`, `file_name`, `file_size`," +
				"`file_addr`, `status`) values (?,?,?,?,1)")


		if err!= nil{
			fmt.Println("failed to prepared statement err:" + err.Error())
			return false
	}

	defer stmt.Close()

	//exec command sql

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)

	if err != nil{
		fmt.Println(err.Error())
		return false
	}

	//if rf <=0 represent OK and return true
	if rf, err := ret.RowsAffected(); nil == err{
		if rf <= 0{
			fmt.Print("File with hash: %s", filehash)
		}
		return true
	}

	return false


}


type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

func GetFileMeta(filehash string)(* TableFile, error){
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_addr,file_name,file_size from tbl_file where file_sha1=? and status=1 limit 1")

	if err != nil{
		fmt.Println(err.Error())
		return nil, err
	}

	defer stmt.Close()

	title := TableFile{}

	err = stmt.QueryRow(filehash).Scan(
		&title.FileHash, &title.FileAddr, &title.FileName, &title.FileSize)

	if err!=nil{
		fmt.Println(err.Error())
		return nil, err
	}
	return &title, nil
}
