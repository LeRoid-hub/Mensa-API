package cache

import (
	"errors"

	"github.com/LeRoid-hub/Mensa-API/models"
)

var Cache = make(map[string]models.CacheItem)

func HasCacheData(key string) bool {
	_, ok := Cache[key]
	return ok
}

func GetCacheData(key string) (models.Mensa, error) {
	Item, ok := Cache[key]
	if !ok {
		return models.Mensa{}, errors.New("no data in cache")
	}
	return Item.GetData()
}

func SetCacheData(key string, data models.Mensa, lifetime ...int64) {
	Item, ok := Cache[key]
	if !ok {
		Item = models.CacheItem{}
	}
	Item.SetData(data, lifetime...)
	Cache[key] = Item
}
