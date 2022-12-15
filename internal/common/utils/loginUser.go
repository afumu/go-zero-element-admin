package utils

import (
	"context"
	"encoding/json"
)

func GetLoginUserId(ctx context.Context) int64 {
	userIdAny := ctx.Value("userId")
	userIdNumber := userIdAny.(json.Number)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return 0
	}
	return userId
}
