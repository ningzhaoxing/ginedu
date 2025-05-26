package routers

import (
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter(e *gin.Engine) {
	e.GET("/sse", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache", "no-cache")
		c.Header("Connection", "keep-alive")

		for i := 0; i < 10; i++ {
			select {
			case <-c.Writer.CloseNotify():
				fmt.Println("关闭了")
				return
			default:
				err := sse.Encode(c.Writer, sse.Event{Data: fmt.Sprintf("data %d", i)})
				if err != nil {
					return
				}
				c.Writer.Flush()
				time.Sleep(time.Second * 3)
			}
		}
	})
	e.LoadHTMLGlob("internal/sse/view/*")
	e.GET("/sseHtml", func(c *gin.Context) {

		c.HTML(http.StatusOK, "sse.html", gin.H{})
	})
}
