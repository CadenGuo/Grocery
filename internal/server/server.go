package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"grocery/internal/server/api"
)

type Config struct {
	Addr  string
	Debug bool
}

func ListenAndServe(ctx context.Context, conf Config, handler *api.Handler) (err error) {
	if handler == nil {
		return errors.New("invalid handler")
	}

	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := &http.Server{
		Addr:    conf.Addr,
		Handler: routerEngine(handler),
	}

	var g errgroup.Group
	g.Go(func() error {
		<-ctx.Done()
		timeout := time.Duration(10) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		return server.Shutdown(ctx)
	})
	g.Go(func() error {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	return g.Wait()
}
