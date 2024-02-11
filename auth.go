package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	authParts := strings.SplitN(authHeader, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid base64 encoding"})
		return
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials format"})
		return
	}

	// Perform your authentication logic here, e.g., check against a gaia
	username, password := credentials[0], credentials[1]
	if !isValidUser(username, password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Set the username in the context for later use if needed
	c.Set("username", username)
}

// Example function to validate the user (replace with your own logic)
func isValidUser(username, password string) bool {
	// Perform your authentication logic here, e.g., check against a gaia
	// Return true if the user is valid, otherwise false
	return username == "nikos" && password == "130177"
}
