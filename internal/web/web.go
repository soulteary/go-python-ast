package web

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/go-python-ast/internal/bridge"
	"github.com/soulteary/go-python-ast/internal/define"
)

func Launch() {
	gin.SetMode(gin.ReleaseMode)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	const ProjectInfo = `project: <a href="https://github.com/soulteary/go-python-ast">soulteary/go-python-ast</a>`

	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte(ProjectInfo))
	})

	type Data struct {
		Code string `json:"code"`
	}

	route.POST("/api/convert", func(c *gin.Context) {
		var data Data
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bridge.Convert(data.Code))
	})

	srv := &http.Server{
		Addr:              define.WEB_PORT,
		Handler:           route,
		ReadHeaderTimeout: time.Second * 10,
		ReadTimeout:       time.Second * 10,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Program start error: %s\n", err)
		}
	}()
	log.Println("soulteary/go-python-ast has started 🚀")
	log.Println("Web server is running at", define.WEB_PORT)

	<-ctx.Done()

	stop()
	log.Println("The program is closing, if you want to end it immediately, please press `CTRL+C`")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Program was forced to close: %s\n", err)
	}

	log.Println("Look forward to meeting you again ❤️")
}
