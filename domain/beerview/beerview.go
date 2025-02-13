package beerview

import (
	"context"
	"github.com/MikelSot/amaris-beer/model"
	"time"
)

type UseCase interface {
	Increment(ctx context.Context, beerID uint) (int64, error)
	ResetViewCounter(ctx context.Context, beerID uint) error
	IsHighDemandReached(ctx context.Context, beerID uint) bool

	PublishHighDemand(ctx context.Context, beerID uint) error
}

type RedisService interface {
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error
	Incr(ctx context.Context, key string) (int64, error)
	Get(ctx context.Context, key string) (string, error)
}

type StreamService interface {
	Publish(ctx context.Context, event model.Event, body []byte) error
}
