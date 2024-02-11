package main

import (
	"apigo/gaia"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getPersons Called"})
}

func getPersonById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "getPersonById " + id + " Called"})
}

func addPerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "addPerson Called"})
}

func updatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updatePerson Called"})
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deletePerson " + id + " Called"})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "options Called"})
}

/*
Albums responds with the list of all albums as JSON.
c.JSON(http.StatusOK, gin.H{"message": "getPersons Called"})
*/
// albums slice to seed record album data.
var albums = []gaia.Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	if err := c.BindJSON(albums); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Albums not found"})
	}
	c.IndentedJSON(http.StatusOK, albums)
}
func getAlbumById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Handle the error (e.g., invalid input for id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	albumData := albums[id]
	c.IndentedJSON(http.StatusOK, albumData)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum gaia.Album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}
