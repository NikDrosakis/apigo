package gaia

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func Globsquery(tableName string) ([]*Globs, error) {
	slice := make([]*Globs, 0)
	db, err := sql.Open("mysql", os.Getenv("MARIA_CONNECTOR"))
	if err != nil {
		return nil, err
	}
	defer db.Close()
	//2nd
	query := "SELECT name,en FROM " + tableName
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	for results.Next() {
		var name string
		var en string
		err := results.Scan(&name, &en)
		if err != nil {
			log.Fatal(err)
		}
		//3rd
		item := &Globs{
			Name: name,
			En:   en,
		}
		slice = append(slice, item)
	}
	return slice, nil
}
