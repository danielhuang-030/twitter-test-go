package model

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const PrefixBlacklistToken string = "blacklist_token"

var ctx = context.Background()

func FindBlacklistToken(token string) (result bool, err error) {
	result = true
	_, err = rdb.Get(ctx, getBlacklistTokenKey(token)).Result()
	if err != nil {
		result = false
		if err == redis.Nil {
			// err = errors.New("Token does not exists")
			err = nil
		}
		return
	}
	return
}

func CreateBlacklistToken(token string, expireAt int64) (err error) {
	if err = rdb.Set(ctx, getBlacklistTokenKey(token), token, time.Until(time.Unix(expireAt, 0))).Err(); err != nil {
		return
	}
	return
}

func getBlacklistTokenKey(token string) string {
	return fmt.Sprintf("%s_%s", PrefixBlacklistToken, token)
}
