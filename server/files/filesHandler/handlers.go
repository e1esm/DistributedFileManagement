package filesHandler

import (
	"DistributedFileManagement/server/config"
	"DistributedFileManagement/server/files/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
)

type Handler interface {
	SaveFile(c *gin.Context)
	GetFiles(c *gin.Context)
}

type FileHandler struct {
	service *service.FileService
}

func NewFileHandler(fileService *service.FileService) *FileHandler {
	return &FileHandler{fileService}
}

func SetFilesHandlers(engine *gin.Engine, handler Handler) {
	engine.GET("/files", handler.GetFiles)
	engine.POST("/files", handler.SaveFile)

}

func (h *FileHandler) SaveFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")

	if err != nil {
		config.Logger.Error("Couldn't have fetched a file", zap.String("err", err.Error()))
		c.JSON(409, "Invalid file")
	}
	file, _ := fileHeader.Open()
	defer file.Close()
	bytesReceived, err := io.ReadAll(file)
	if err != nil {
		config.Logger.Error("Couldn't have read a file", zap.String("err", err.Error()))
		c.JSON(409, "Error while reading a file")
	}

	c.JSON(200, "Fetched file")
}

func (h *FileHandler) GetFiles(c *gin.Context) {
	c.JSON(200, "all files")
}
