package gaia

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Postquery(tableName string) ([]*Post, error) {
	slice := make([]*Post, 0)
	db, err := sql.Open("mysql", "root:n130177@tcp(localhost:9906)/ga_240204?parseTime=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	//2nd
	query := "SELECT id,uid,uri,img,title,subtitle,excerpt,content FROM " + tableName
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	for results.Next() {
		var id int
		var uid int
		var uri string
		var img string
		var title string
		var subtitle string
		var excerpt string
		var content string
		err := results.Scan(&id, &uid, &uri, &img, &title, &subtitle, &excerpt, &content)
		if err != nil {
			log.Fatal(err)
		}
		//3rd
		item := &Post{
			ID:       id,
			UID:      uid,
			Uri:      uri,
			Img:      img,
			Title:    title,
			Subtitle: subtitle,
			Excerpt:  excerpt,
			Content:  content,
		}
		slice = append(slice, item)
	}
	return slice, nil
}
