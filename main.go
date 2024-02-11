package main

import (
	"apigo/gaia"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// API v1
// router := r.Group("/api/v1")
func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Update with your allowed origins
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"} // Include other headers as needed
	router.Use(cors.New(config))
	//middleware for basic auth
	router.Use(authMiddleware)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/:table", func(c *gin.Context) {
			tableName := c.Param("table")
			var err error
			Data, err := gaia.Postquery(tableName)
			//error
			if err != nil {
				// Handle the error and respond with an error message
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Respond with JSON
			c.IndentedJSON(http.StatusOK, Data)
		})
		v1.GET("/albums", getAlbums)
		v1.GET("/albums/:id", getAlbumById)
		v1.POST("/albums", postAlbums)
		v1.GET("/person", getPersons)
		v1.GET("/person/:id", getPersonById)
		v1.POST("/person", addPerson)
		v1.PUT("/person/:id", updatePerson)
		v1.DELETE("/person/:id", deletePerson)
		v1.OPTIONS("/person", options)
	}
	router.Run(":8082")
}
