package cache

import (
	"HOPE-backend/config"
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type inmem struct {
	cache *cache.Cache
}

func New(cfg config.CacheConfig) Cache {
	var (
		ttl   time.Duration = 600
		purge time.Duration = 600
	)

	if cfg.Inmem.TtlInSecond != 0 {
		ttl = cfg.Inmem.TtlInSecond
	}

	if cfg.Inmem.PurgeTimeInSecond != 0 {
		purge = cfg.Inmem.PurgeTimeInSecond
	}

	return &inmem{cache: cache.New(ttl*time.Second, purge*time.Second)}
}

func (i *inmem) Set(ctx context.Context, key string, value interface{}, expiryInSec int64) {
	i.cache.Set(key, value, time.Duration(expiryInSec)*time.Second)
}

func (i *inmem) Get(ctx context.Context, key string) (interface{}, bool) {
	return i.cache.Get(key)
}

func (i *inmem) Increment(ctx context.Context, key string, n int64) error {
	if err := i.cache.Increment(key, n); err != nil {
		return fmt.Errorf("[Cache.Inmem.Increment][990001] Failed: %v", err)
	}

	return nil
}
