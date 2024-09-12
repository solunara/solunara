package redislimit

import "context"

type Limiter interface {
	// key: limiting object
	// return true is limited
	Limit(ctx context.Context, key string) (bool, error)
}
