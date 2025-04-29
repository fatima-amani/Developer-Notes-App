package controller

import (
	"developer-notes/config"
	"developer-notes/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET /notes
func GetNotes(c *gin.Context) {
	var notes []models.Note
	result := config.DB.Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}

// GET /notes/:id
func GetNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note
	result := config.DB.First(&note, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}

// POST /notes
func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	note.ID = uuid.NewString()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	if err := config.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create note"})
		return
	}
	c.JSON(http.StatusCreated, note)
}

// PUT /notes/:id
func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note
	if err := config.DB.First(&note, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var updatedData models.Note
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	note.Title = updatedData.Title
	note.Content = updatedData.Content
	note.Tags = updatedData.Tags
	note.UpdatedAt = time.Now()

	if err := config.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update note"})
		return
	}
	c.JSON(http.StatusOK, note)
}

// DELETE /notes/:id
func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Note{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
