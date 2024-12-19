package cache

import (
	"EurikaOrmanel/up-charter/internal/models"
	"encoding/json"
	"log"
	"time"
)

func (cacheConfig CacheConfig) GetTop100Chart() []models.Top100Chart {
	key := "top_100"

	result, err := cacheConfig.RedisClient.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
		return []models.Top100Chart{}
	}
	resp := []models.Top100Chart{}
	err = json.Unmarshal([]byte(result), &resp)
	if err != nil {
		log.Println(err)
		return []models.Top100Chart{}
	}
	return resp
}

func (cacheConfig CacheConfig) SetTop100Chart(chart []models.Top100Chart) error {
	key := "top_100"
	jsonByte, err := json.Marshal(chart)
	if err != nil {
		return err
	}
	return cacheConfig.RedisClient.Set(ctx, key, jsonByte, time.Hour*24).Err()
}
