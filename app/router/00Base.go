package router

import "github.com/gin-gonic/gin"

func RegisteredRoutes(r *gin.Engine) {
	// 本地路径相关
	registeredLocalPath(r.Group("/localPath"))
}
