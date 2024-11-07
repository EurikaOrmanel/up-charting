package cache

import (
	"fmt"
	"strings"
	"time"
)

func (cacheConfig CacheConfig) SetRefreshToken(authId string, token string) error {
	key := fmt.Sprintf("refresh_token:%s", authId)
	value := "value:" + token
	err := cacheConfig.Client.Set(ctx, key, value, time.Minute*15).Err()
	if err != nil {
		return err
	}
	return cacheConfig.Client.Set(ctx, value, key, time.Minute*5).Err()

}

func (cacheConfig CacheConfig) FindUserIdByToken(refreshToken string) (string, error) {
	value := "value:" + refreshToken
	foundTokens, err := cacheConfig.Client.Get(ctx, value).Result()
	if err != nil {
		return "", err
	}
	fmt.Println(foundTokens)
	valueFromKey := strings.Split(foundTokens, ":")[1]
	return valueFromKey, nil

}

func (cacheConfig CacheConfig) DeleteTokenCache(authId string) error {
	key := fmt.Sprintf("refresh_token:%s", authId)
	foundToken, err := cacheConfig.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = cacheConfig.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return cacheConfig.Client.Del(ctx, foundToken).Err()
}
