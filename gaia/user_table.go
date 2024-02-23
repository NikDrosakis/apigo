package gaia

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"os"
)

func Userquery(c *gin.Context) {
	// Retrieve id from the URL parameters
	id := c.Param("id")
	var err1 error
	err1 = godotenv.Load(".env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	slice := make([]*User, 0)
	db, err := sql.Open("mysql", os.Getenv("MARIA_CONNECTOR"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()
	//2nd
	var results *sql.Rows
	query := "SELECT id,url,img,name,pass,firstname,lastname FROM user"
	if id != "" {
		// If id is not empty, add a WHERE clause to filter by id
		query += " WHERE id = ?"
		results, err = db.Query(query, id)
	}else{
		results, err = db.Query(query)
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer results.Close()
	for results.Next() {
		var id int
		var url sql.NullString
		var img sql.NullString
		var name string
		var pass string
		var firstname string
		var lastname string
		err := results.Scan(&id,&url,&img, &name, &pass, &firstname, &lastname)
		if err != nil {
			// Handle the error by sending an error response
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//3rd
		item := &User{
			ID:        id,
			Url:       url,
			Img:       img,
			Name:      name,
			Pass:      pass,
			Firstname: firstname,
			Lastname:  lastname,
		}
		slice = append(slice, item)
	}
	c.IndentedJSON(http.StatusOK, slice)
}