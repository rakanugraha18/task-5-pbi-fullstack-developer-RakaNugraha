package controllers

import (
	"net/http"
	"task_5_pbi_btpns_RakaNugraha/database"
	"task_5_pbi_btpns_RakaNugraha/middlewares"
	"task_5_pbi_btpns_RakaNugraha/models"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	userID, err := middlewares.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID from token"})
		return
	}

	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set UserID pada foto
	photo.UserID = int(userID)

	// Simpan foto ke database
	if err := database.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo created successfully"})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo

	// Ambil semua foto dari database beserta informasi pengguna terkait
	if err := database.DB.Preload("User").Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	photoID := c.Param("photoId")

	// Cari foto berdasarkan ID
	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Periksa izin pengguna saat ini untuk memperbarui foto
	userIDFromToken := getUserIDFromToken(c)
	if int(photo.UserID) != int(userIDFromToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this photo"})
		return
	}

	// Bind data baru ke photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func DeletePhoto(c *gin.Context) {
	// Dapatkan ID pengguna dari token JWT
	userIDFromToken, err := middlewares.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Dapatkan ID foto dari URL
	photoID := c.Param("photoId")

	// Cari foto berdasarkan ID
	var photo models.Photo
	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Periksa apakah pengguna adalah pemilik foto
	if int(photo.UserID) != int(userIDFromToken) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this photo"})
		return
	}

	// Hapus foto dari database
	if err := database.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

func getUserIDFromToken(c *gin.Context) uint {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(uint)
}
