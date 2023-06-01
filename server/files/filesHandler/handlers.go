package filesHandler

import (
	"github.com/gin-gonic/gin"
)

func SetFilesHandlers(engine *gin.Engine) {
	engine.GET("/files", getFiles)
	engine.POST("/files", saveFile)

}

func saveFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		//zap.Error("Error while saving file", zap.String("Error", err.Error()))
	}
}

func getFiles(c *gin.Context) {
	c.JSON(200, "all files")
}
