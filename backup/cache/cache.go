package cache

import (
	"mensa/types"
)

func NewCache() *types.Cache {
	cache := types.Cache{}
	return &cache
}

func SetCacheData(cache *types.Cache, key string, data string, lifetime ...int64) {
	cache.SetData(key, data, lifetime...)
}

func GetCacheData(cache *types.Cache) string {
	return cache.GetData()
}
