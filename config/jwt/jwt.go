package jwt

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"github.com/joker-star-l/dousheng_common/config/log"
	common "github.com/joker-star-l/dousheng_common/entity"
	util_json "github.com/joker-star-l/dousheng_common/util/json"
	"strconv"
	"time"
)

const (
	KeyIdentity = "identity"
	KeyToken    = "token"
	KeyData     = "data"
)

var Middleware *jwt.HertzJWTMiddleware

func init() {
	var err error
	Middleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:   "service_demo",
		Key:     []byte("service_demo_key"),
		Timeout: time.Hour * 24 * 7,
		//MaxRefresh:                  time.Hour * 24 * 7,
		TokenLookup:                 fmt.Sprintf("header: %s, query: %s, param: %s, form: %s", KeyToken, KeyToken, KeyToken, KeyToken),
		WithoutDefaultTokenHeadName: true,
		// 向 token 中添加信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{KeyData: data}
		},
		// 解析 token
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[KeyData]
		},
		// 认证成功 权限校验
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			log.Slog.Infof("auth success: %v", util_json.Str(data))
			return true
		},
		// 认证失败
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			log.Slog.Infof("auth error: %v", message)
			c.JSON(consts.StatusOK, common.ErrorResponse("未登录"))
		},
	})
	if err != nil {
		log.Slog.Panicf("HertzJWTMiddleware init error: %v", err.Error())
	}
}

func ParseAndGetUserId(c context.Context, ctx *app.RequestContext) int64 {
	userId := int64(0)
	claims, err := Middleware.GetClaimsFromJWT(c, ctx)
	if err == nil && claims != nil {
		userId, _ = strconv.ParseInt(claims[KeyData].(map[string]any)["id"].(string), 10, 0)
	}
	return userId
}

func GetUserId(ctx *app.RequestContext) int64 {
	tokenUser, _ := ctx.Get(KeyIdentity)
	userId, _ := strconv.ParseInt(tokenUser.(map[string]any)["id"].(string), 10, 0)
	return userId
}
