package routers

import (
	"tugas13-bioskop/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.POST("/bioskop", controllers.CreateBioskop)

	return r
}
