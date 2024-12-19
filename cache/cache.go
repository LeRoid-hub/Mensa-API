package cache

import (
	"errors"

	"github.com/LeRoid-hub/Mensa-API/models"
)

var Cache = make(map[string]models.CacheItem[interface{}])

func HasCacheData(key string) bool {
	data, ok := Cache[key]
	if !ok {
		return false
	}
	if data.IsExpired() {
		delete(Cache, key)
		return false
	}
	return ok
}

func GetCacheData(key string) (interface{}, error) {
	Item, ok := Cache[key]
	if !ok {
		return models.Mensa{}, errors.New("no data in cache")
	}
	return Item.GetData()
}

func SetCacheData(key string, data interface{}, lifetime ...int64) {
	Item, ok := Cache[key]
	if !ok {
		Item = models.CacheItem[interface{}]{}
	}
	Item.SetData(data, lifetime...)
	Cache[key] = Item
}
