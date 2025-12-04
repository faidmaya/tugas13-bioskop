package routers

import (
	"tugas13-bioskop/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	// POST
	r.POST("/bioskop", controllers.CreateBioskop)

	// GET ALL
	r.GET("/bioskop", controllers.GetAllBioskop)

	// GET BY ID
	r.GET("/bioskop/:id", controllers.GetBioskopByID)

	// UPDATE
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)

	// DELETE
	r.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	return r
}
