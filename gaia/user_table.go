package gaia

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func Userquery(tableName string) ([]*User, error) {
	slice := make([]*User, 0)
	db, err := sql.Open("mysql", os.Getenv("MARIA_CONNECTOR"))
	if err != nil {
		return nil, err
	}
	defer db.Close()
	//2nd
	query := "SELECT id,uid,url,img,name,pass,firstname,lastname FROM " + tableName
	results, err := db.Query(query, os.Getenv("MARIA_CONNECTOR"))
	if err != nil {
		return nil, err
	}
	defer results.Close()
	for results.Next() {
		var id int
		var uid int
		var url string
		var img string
		var name string
		var pass string
		var firstname string
		var lastname string
		err := results.Scan(&id, &uid, &url, &img, &name, &pass, &firstname, &lastname)
		if err != nil {
			log.Fatal(err)
		}
		//3rd
		item := &User{
			ID:        id,
			UID:       uid,
			Url:       url,
			Img:       img,
			Name:      name,
			Pass:      pass,
			Firstname: firstname,
			Lastname:  lastname,
		}
		slice = append(slice, item)
	}
	return slice, nil
}
