package main

import (
	"DistributedFileManagement/server/files/filesHandler"
	"DistributedFileManagement/server/files/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.MaxMultipartMemory = 12 << 20
	r.Use()
	r.Use(gin.Recovery())
	fileHandler := filesHandler.NewFileHandler(service.NewFileService())
	filesHandler.SetFilesHandlers(r, fileHandler)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
