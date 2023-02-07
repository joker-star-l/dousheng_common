package util_hertz

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/joker-star-l/dousheng_common/config/log"
	common "github.com/joker-star-l/dousheng_common/entity"
	"strconv"
	"time"
)

func InitServer(port int) *server.Hertz {
	h := server.New(server.WithHostPorts(":" + strconv.Itoa(port)))
	// 请求信息
	h.Use(func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		latency := time.Now().Sub(start)
		log.Slog.Infof("cost: %v, request: %v", latency, ctx.URI())
	})
	// 全局 panic 处理
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(func(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
		log.Slog.Errorf("Recovery, err: %v\nstack: %s", err, stack)
		ctx.JSON(consts.StatusInternalServerError, common.ErrorResponse("系统错误"))
	})))
	return h
}

// UseCors 跨域
func UseCors(h *server.Hertz, allowOrigins []string) {
	if nil != allowOrigins {
		config := cors.Config{
			AllowOrigins:     allowOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			MaxAge:           12 * time.Hour,
			AllowWildcard:    true,
		}
		h.Use(cors.New(config))
	}
}
