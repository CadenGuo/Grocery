package api

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"grocery/internal/db"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	modeLive = "live" //nolint:unused,deadcode,varcheck
	modeTest = "test" //nolint:unused,deadcode,varcheck
)

type Handler struct {
	wg        sync.WaitGroup
	dbManager *db.Manager
	debug     bool // not run in live env
}

func NewHandler(
	dbManager *db.Manager,
	debug bool,
) (*Handler, error) {
	h := Handler{
		debug:     debug,
		dbManager: dbManager,
	}
	return &h, nil
}

func (h *Handler) WaitAndClose() {
	h.wg.Wait()
}

func (h *Handler) LogRequest(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		log.Info().Str("request", ctx.Request.URL.String()).Msg("request_log")
	} else if ctx.Request.Method == http.MethodPost || ctx.Request.Method == http.MethodPut {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		log.Info().Bytes("request", body).Msg("request_log")
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	}

}

type response struct {
	Error    bool        `json:"error"`
	ErrorMsg string      `json:"error_msg"`
	Data     interface{} `json:"data"`
}

func failedResponse(data interface{}, errorMsg string) response {
	return response{
		Error:    true,
		ErrorMsg: errorMsg,
		Data:     data,
	}
}

func succeededResponse(data interface{}) response {
	return response{
		Error:    false,
		ErrorMsg: "",
		Data:     data,
	}
}
