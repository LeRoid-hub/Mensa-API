package models

import (
	"errors"
	"time"
)

type CacheItem struct {
	data        Mensa
	lastUpdated time.Time
	lifetime    int64
}

func (c *CacheItem) SetData(data Mensa, lifetime ...int64) {
	if len(lifetime) > 0 {
		c.lifetime = lifetime[0]
	} else {
		c.lifetime = 60
	}

	c.data = data
	c.lastUpdated = time.Now()
}

func (c *CacheItem) GetData() (Mensa, error) {
	if time.Now().Unix()-c.lastUpdated.Unix() > c.lifetime {
		return Mensa{}, errors.New("cache expired")
	}

	if c.lastUpdated.IsZero() {
		return Mensa{}, errors.New("no data in cache")
	}

	return c.data, nil
}
