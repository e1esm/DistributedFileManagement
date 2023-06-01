package main

import (
	"DistributedFileManagement/server/files/filesHandler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.MaxMultipartMemory = 12 << 20
	r.Use(gin.Recovery())
	filesHandler.SetFilesHandlers(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
