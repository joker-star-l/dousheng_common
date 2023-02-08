package util_kitex

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/joker-star-l/dousheng_common/config/log"
)

func Log(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		err := next(ctx, req, resp)
		log.Slog.Infof("kitex request: %v, response: %v", req, resp)
		return err
	}
}
