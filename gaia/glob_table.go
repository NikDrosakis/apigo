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

func Globsquery(c *gin.Context) {
	name := c.Param("name")
	var err1 error
	err1 = godotenv.Load(".env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	slice := make([]*Globs, 0)
	db, err := sql.Open("mysql", os.Getenv("MARIA_CONNECTOR"))
	if err != nil {

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()
	//2nd
	var results *sql.Rows
	query := "SELECT name,en FROM globs"
	if name != "" {
		// If id is not empty, add a WHERE clause to filter by id
		query += " WHERE name = ?"
		results, err = db.Query(query, name)
	}else {
		results, err = db.Query(query)
	}
	if err != nil {

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	defer results.Close()
	for results.Next() {
		var name string
		var en string
		err := results.Scan(&name, &en)
		if err != nil {
			// Handle the error by sending an error response
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//3rd
		item := &Globs{
			Name: name,
			En:   en,
		}
		slice = append(slice, item)
	}
	c.IndentedJSON(http.StatusOK, slice)
}
