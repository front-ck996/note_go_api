package router

import (
	"github.com/gin-gonic/gin"
	"note_go_api/app/controller"
	"note_go_api/app/controller/localPath"
)

func registeredLocalPath(r *gin.RouterGroup) {
	r.GET("/tree", controller.Handle(localPath.Tree))
}
