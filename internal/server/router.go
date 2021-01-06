package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"grocery/internal/server/api"
)

func routerEngine(h *api.Handler) *gin.Engine {
	r := gin.New()
	r.Use(logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")

	api.POST("/create_grocery_item", h.GroceryItem)
	api.PUT("/update_grocery_item", h.GroceryItem)
	api.GET("/list_grocery_item", h.GroceryItem)
	api.POST("/create_dishes", h.Dishes)
	api.GET("/list_dishes", h.Dishes)
	return r
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}

		end := time.Now()
		latency := end.Sub(start)

		dumpLogger := log.With().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("client-ip", c.ClientIP()).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Int("body-size", c.Writer.Size()).
			Str("user-agent", c.Request.UserAgent()).
			Logger()
		if len(c.Errors) == 0 {
			dumpLogger.Info().
				Msg(msg)
		} else {
			dumpLogger.Error().Strs("errors", c.Errors.Errors()).
				Msg(msg)
		}
	}
}
