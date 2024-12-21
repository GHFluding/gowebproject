package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	webServer()
}

// fjsj
func webServer() error {
	//new gin-gonic
	r := gin.Default()

	r.LoadHTMLGlob("HTML/*.html")

	r.GET("/", handlerIndex)

	err := r.Run()
	if err != nil {
		return err
	}
	return nil
}

func handlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
