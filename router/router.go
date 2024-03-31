package router

import (
	"task_5_pbi_btpns_RakaNugraha/controllers"
	"task_5_pbi_btpns_RakaNugraha/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Membuat instance Gin router
	r := gin.Default()

	// Endpoint untuk registrasi pengguna
	r.POST("/users/register", controllers.Register)

	// Endpoint untuk login pengguna
	r.POST("/users/login", controllers.Login)

	// Endpoint untuk meng-update pengguna
	r.PUT("/users/:userId", middlewares.AuthMiddleware(), controllers.UpdateUser)

	// Endpoint untuk menghapus pengguna
	r.DELETE("/users/:userId", middlewares.AuthMiddleware(), controllers.DeleteUser)

	// Endpoint untuk membuat foto baru
	r.POST("/photos", middlewares.AuthMiddleware(), controllers.CreatePhoto)

	// Endpoint untuk mendapatkan daftar foto
	r.GET("/photos", controllers.GetPhotos)

	// Endpoint untuk meng-update foto
	r.PUT("/photos/:photoId", middlewares.AuthMiddleware(), controllers.UpdatePhoto)

	// Endpoint untuk menghapus foto
	r.DELETE("/photos/:photoId", middlewares.AuthMiddleware(), controllers.DeletePhoto)

	return r
}
