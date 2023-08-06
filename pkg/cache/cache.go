package cache

import "context"

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiryInSec int64)
	Get(ctx context.Context, key string) (interface{}, bool)
	Increment(ctx context.Context, key string, n int64) error
}
