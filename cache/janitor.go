package cache

import (
	"context"
	"time"

	"github.com/soverenio/vanilla/throw"

	"github.com/soverenio/vanilla/promise"
)

type janitor[K comparable, V any] struct {
	interval time.Duration
	cache    *cache[K, V]
	task     *promise.Simple[struct{}]
}

func (j *janitor[K, V]) Start(ctx context.Context) {
	if j.task != nil {
		panic(throw.IllegalState())
	} else if j.interval <= 0 {
		return
	}

	j.task = promise.NewSimpleStarted[struct{}](ctx, j.loop)
}

func (j *janitor[K, V]) loop(ctx context.Context) (struct{}, error) {
	ticker := time.NewTicker(j.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return struct{}{}, nil
		case <-ticker.C:
		}

		j.cache.DeleteExpired()
	}
}

func (j *janitor[K, V]) Stop() {
	if j.task == nil {
		return
	}

	j.task.Cancel()
}

func newJanitor[K comparable, V any](c *cache[K, V], ci time.Duration) *janitor[K, V] {
	return &janitor[K, V]{
		interval: ci,
		cache:    c,
	}
}
