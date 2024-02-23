package main

import (
	"apigo/gaia"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
		"os"
		"log"
	"github.com/joho/godotenv"
)
// API v1
// router := r.Group("/api/v1")
func main() {
    var err1, err3 error
	err1 = godotenv.Load(".env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://nikosdrosakis.gr"} // Update with your allowed origins
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"} // Include other headers as needed
	router.Use(cors.New(config))
	//middleware for basic auth
	router.Use(authMiddleware)

	v1 := router.Group("/apigo")
        v1.GET("/albums", getAlbums)
        v1.GET("/albums/:id", getAlbumById)
        v1.POST("/albums", postAlbums)
        v1.GET("/person", getPersons)
        v1.GET("/person/:id", getPersonById)
        v1.POST("/person", addPerson)
        v1.PUT("/person/:id", updatePerson)
        v1.DELETE("/person/:id", deletePerson)
        v1.OPTIONS("/person", options)
        v1.GET("/user", gaia.Userquery)
        v1.GET("/user/:id", gaia.Userquery)
        v1.GET("/post", gaia.Postquery)
        v1.GET("/post/:id", gaia.Postquery)
		v1.GET("/globs", gaia.Globsquery)
        v1.GET("/globs/:name", gaia.Globsquery)
	// Get certificate and private key filenames from environment variables
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")
	
	err3 = router.RunTLS(":8082",certFile,keyFile)
//	router.Run(":8082")
if err3 != nil {
    log.Fatal("Failed to start server with TLS: ", err3)
}
}
