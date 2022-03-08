package main

import (
	"flag"
	"net/http"
	"path"
	"time"

	"go.b8s.dev/static/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var configPath = flag.String("config", "config.yaml", "Path to the config file.")

func main() {
	flag.Parse()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	r := gin.New()
	r.SetTrustedProxies(conf.TrustedProxies)
	r.Use(gin.Recovery(), requestLogger(logger))
	r.Use(customHeaderInjector(conf.Headers))

	r.Static("/", conf.Webroot)

	r.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
		c.File(path.Join(conf.Webroot, "error.html"))
	})

	logger.Info("Server started on :8080.")
	r.Run(":8080")
}

func requestLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestStart := time.Now()
		c.Next()

		logger.Info("Request Processed",
			zap.String("method", c.Request.Method),
			zap.Int("status_code", c.Writer.Status()),
			zap.String("path", c.Request.URL.Path),
			zap.Duration("latency", time.Since(requestStart)),
		)
	}
}

func customHeaderInjector(headers []config.ResponseHeader) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, h := range headers {
			c.Header(h.Name, h.Value)
		}
	}
}
