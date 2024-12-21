package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//new gin-gonic
	r := gin.Default()

	r.LoadHTMLGlob("HTML/*.html")

	r.GET("/", handlerIndex)

	r.Run()

}
func handlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
