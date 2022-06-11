package router

import "github.com/gin-gonic/gin"

func RegisteredRoutes(r *gin.Engine) {
	rApiV1 := r.Group("/api/v1")

	// 本地路径相关
	registeredLocalPath(rApiV1.Group("/localPath"))
}
