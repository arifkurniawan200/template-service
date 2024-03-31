package cache

import "context"

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl int) error
	Delete(ctx context.Context, key string) error
}

type CacheManager struct {
	client Cache
}

func (c *CacheManager) Set(ctx context.Context, key string, value string, ttl int) error {
	return c.client.Set(ctx, key, value, ttl)
}

func (c *CacheManager) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key)
}

func (c *CacheManager) Delete(ctx context.Context, key string) error {
	return c.client.Delete(ctx, key)
}
