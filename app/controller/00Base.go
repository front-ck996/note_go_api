package controller

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

//返回JSON
func (c *Context) RJson(data interface{}) {
	c.JSON(200, data)
}
