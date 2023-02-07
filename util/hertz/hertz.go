package util_hertz

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/jwt"
	"github.com/joker-star-l/dousheng_common/config/log"
	common "github.com/joker-star-l/dousheng_common/entity"
	"strconv"
	"time"
)

const (
	KeyIdentity = "identity"
	KeyToken    = "token"
	KeyData     = "data"
)

func InitServer(port int) *server.Hertz {
	h := server.New(server.WithHostPorts(":" + strconv.Itoa(port)))
	// 请求信息
	h.Use(func(c context.Context, ctx *app.RequestContext) {
		log.Slog.Infof("request: %v", ctx.URI())
	})
	// 全局 panic 处理
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(func(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
		log.Slog.Errorf("%v, [Recovery] err=%v\nstack=%s", c, err, stack)
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

func DefaultJwtMiddleware() *jwt.HertzJWTMiddleware {
	jwtMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:                       "service_demo",
		Key:                         []byte("service_demo_key"),
		Timeout:                     time.Hour,
		MaxRefresh:                  time.Hour * 24 * 7,
		TokenLookup:                 fmt.Sprintf("header: %s, query: %s, param: %s, form: %s", KeyToken, KeyToken, KeyToken, KeyToken),
		WithoutDefaultTokenHeadName: true,
		// 登录认证
		//		Authenticator: nil,
		// 登录时向 token 中添加字段
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{KeyData: data}
		},
		// 登录返回值
		//		LoginResponse: nil
		// 解析 token
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[KeyData]
		},
		// 认证成功 权限校验
		//		Authorizator: nil,
		// 认证失败
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			log.Slog.Errorf("auth error: %v", message)
			c.JSON(code, common.ErrorResponse("认证失败"))
		},
	})
	if err != nil {
		log.Slog.Panicf("HertzJWTMiddleware init error: %v", err.Error())
	}
	return jwtMiddleware
}
