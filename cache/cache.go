package cache

import (
	"github.com/LeRoid-hub/Mensa-API/models"
)

var Cache = make(map[string]models.CacheItem)

func HasCacheData(key string) bool {
	_, ok := Cache[key]
	return ok
}

func GetCacheData(key string) string {
	Item, ok := Cache[key]
	if !ok {
		return ""
	}
	return Item.GetData()
}

func SetCacheData(key string, data string, lifetime ...int64) {
	Item, ok := Cache[key]
	if !ok {
		Item = models.CacheItem{}
	}
	Item.SetData(data, lifetime...)
	Cache[key] = Item
}
