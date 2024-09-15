package models

import "time"

type CacheItem struct {
	data        string
	lastUpdated time.Time
	lifetime    int64
}

func (c *CacheItem) SetData(data string, lifetime ...int64) {
	if len(lifetime) > 0 {
		c.lifetime = lifetime[0]
	} else {
		c.lifetime = 60
	}

	c.data = data
	c.lastUpdated = time.Now()
}

func (c *CacheItem) GetData() string {
	if time.Now().Unix()-c.lastUpdated.Unix() > c.lifetime {
		return ""
	}

	if c.data == "" {
		return ""
	}

	return c.data
}
