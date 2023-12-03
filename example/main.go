package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	objimpl "github.com/ntnghiatn/rest-go-gin-api/example/obj/objImpl"
)

func main() {
	cus := objimpl.NewCustomObject()
	cus.Insert("Hello")
	cus.Insert("Ciquan")
	cus.Insert("Kimchon")
	cus.Insert("Conchim")
	cusStr := cus.List()

	fmt.Println("List::::", cusStr)
	// Force log's color
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
