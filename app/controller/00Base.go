package controller

import "github.com/gin-gonic/gin"

type BaseController struct {
	*gin.Context
}

type HandlerFunc func(c *BaseController)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &BaseController{
			c,
		}
		h(ctx)
	}
}

//返回JSON
func (c *BaseController) RJson(data interface{}) {
	c.JSON(200, data)
}
