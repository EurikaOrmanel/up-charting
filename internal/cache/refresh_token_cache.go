package cache

import (
	"fmt"
	"strings"
	"time"
)

func (cacheConfig CacheConfig) SetRefreshToken(authId string, token string) error {
	key := fmt.Sprintf("refresh_token:%s", authId)
	value := "value:" + token
	err := cacheConfig.RedisClient.Set(ctx, key, value, time.Minute*15).Err()
	if err != nil {
		return err
	}
	return cacheConfig.RedisClient.Set(ctx, value, key, time.Minute*5).Err()

}

func (cacheConfig CacheConfig) FindUserIdByToken(refreshToken string) (string, error) {
	value := "value:" + refreshToken
	foundTokens, err := cacheConfig.RedisClient.Get(ctx, value).Result()
	if err != nil {
		return "", err
	}
	fmt.Println(foundTokens)
	valueFromKey := strings.Split(foundTokens, ":")[1]
	return valueFromKey, nil

}

func (cacheConfig CacheConfig) DeleteTokenCache(authId string) error {
	key := fmt.Sprintf("refresh_token:%s", authId)
	foundToken, err := cacheConfig.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = cacheConfig.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return cacheConfig.RedisClient.Del(ctx, foundToken).Err()
}
