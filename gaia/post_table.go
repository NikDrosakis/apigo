package gaia

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
			"github.com/joho/godotenv"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Postquery(c *gin.Context) {
	id := c.Param("id")
	var err1 error
	err1 = godotenv.Load(".env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	slice := make([]*Post, 0)
	db, err := sql.Open("mysql", os.Getenv("MARIA_CONNECTOR"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()
	//2nd
	var results *sql.Rows
	query := "SELECT id,uid,uri,img,title,subtitle,excerpt,content FROM post"
	if id != "" {
		// If id is not empty, add a WHERE clause to filter by id
		query += " WHERE id = ?"
		results, err = db.Query(query, id)
	}else {
		results, err = db.Query(query)
	}
	if err != nil {

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer results.Close()
	for results.Next() {
		var id int
		var uid int
		var uri sql.NullString
		var img sql.NullString
		var title string
		var subtitle sql.NullString
		var excerpt sql.NullString
		var content sql.NullString
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
	c.IndentedJSON(http.StatusOK, slice)
}
