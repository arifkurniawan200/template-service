package cache

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
)

type MemcacheCache struct {
	client *memcache.Client
}

func NewMemcacheCache(client *memcache.Client) *MemcacheCache {
	return &MemcacheCache{client: client}
}

func (c *MemcacheCache) Get(_ context.Context, key string) (string, error) {
	item, err := c.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

func (c *MemcacheCache) Set(_ context.Context, key string, value string, ttl int) error {
	return c.client.Set(&memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: int32(ttl),
	})
}

func (c *MemcacheCache) Delete(_ context.Context, key string) error {
	return c.client.Delete(key)
}
