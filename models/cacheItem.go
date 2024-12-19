package models

import (
	"errors"
	"time"
)

type CacheItem[C any] struct {
	data        C
	lastUpdated time.Time
	lifetime    int64
}

func (c *CacheItem[C]) SetData(data C, lifetime ...int64) {
	if len(lifetime) > 0 {
		c.lifetime = lifetime[0]
	} else {
		c.lifetime = 60
	}

	c.data = data
	c.lastUpdated = time.Now()
}

func (c *CacheItem[C]) GetData() (C, error) {
	if time.Now().Unix()-c.lastUpdated.Unix() > c.lifetime {
		var zeroValue C
		return zeroValue, errors.New("cache expired")
	}

	if c.lastUpdated.IsZero() {
		var zeroValue C
		return zeroValue, errors.New("no data in cache")
	}

	return c.data, nil
}

func (c *CacheItem[C]) IsExpired() bool {
	if time.Now().Unix()-c.lastUpdated.Unix() > c.lifetime {
		return true
	}
	return false
}
