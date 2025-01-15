package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func Get_name(c *gin.Context) {
	file, err := os.Open("resume.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer file.Close()

	// Decode the JSON data
	decoder := json.NewDecoder(file)
	var resume Resume
	err = decoder.Decode(&resume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	// Send the name as a JSON response
	c.JSON(http.StatusOK, gin.H{"name": resume.Name})
}

func Get_contact(c *gin.Context) {
	file, err := os.Open("resume.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer file.Close()

	// Decode the JSON data
	decoder := json.NewDecoder(file)
	var resume Resume
	err = decoder.Decode(&resume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	// Send the name as a JSON response
	c.JSON(http.StatusOK, gin.H{"Cotact": resume.Contact})
}

func Get_experience(c *gin.Context) {
	file, err := os.Open("resume.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer file.Close()

	// Decode the JSON data
	decoder := json.NewDecoder(file)
	var resume Resume
	err = decoder.Decode(&resume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	// Send the name as a JSON response
	c.JSON(http.StatusOK, gin.H{"name": resume.Experience})
}

func Get_education(c *gin.Context) {
	file, err := os.Open("resume.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer file.Close()

	// Decode the JSON data
	decoder := json.NewDecoder(file)
	var resume Resume
	err = decoder.Decode(&resume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	// Send the name as a JSON response
	c.JSON(http.StatusOK, gin.H{"name": resume.Education})
}

// get skills
func Get_skills(c *gin.Context) {
	file, err := os.Open("resume.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer file.Close()

	// Decode the JSON data
	decoder := json.NewDecoder(file)
	var resume Resume
	err = decoder.Decode(&resume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	// Send the name as a JSON response
	c.JSON(http.StatusOK, gin.H{"name": resume.Skills})
}
