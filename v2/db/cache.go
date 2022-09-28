package db

import (
	"HOPE-backend/config"
	"time"

	"github.com/patrickmn/go-cache"
)

func NewInmemCache(cfg config.InmemCacheConfig) *cache.Cache {
	var (
		ttl   time.Duration = 600
		purge time.Duration = 600
	)

	if cfg.TTLInSecond != 0 {
		ttl = cfg.TTLInSecond
	}

	if cfg.PurgeTimeInSecond != 0 {
		purge = cfg.PurgeTimeInSecond
	}

	return cache.New(ttl*time.Second, purge*time.Second)
}
