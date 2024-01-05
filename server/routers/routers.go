package routers

import(
	"github.com/gin-gonic/gin"
	"myfitnesspal-grafana/controllers"
)

func InitRoutes(router *gin.Engine) {
    router.POST("/upload", controllers.UploadCSV)
}

