package util_json

import (
	"encoding/json"
	"github.com/joker-star-l/dousheng_common/config/log"
)

func Str(data any) string {
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Slog.Errorf("%v convert to json error", data)
		return ""
	}
	return string(marshal)
}
