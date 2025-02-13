package beerview

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MikelSot/amaris-beer/model"
)

type BeerView struct {
	cache  RedisService
	stream StreamService

	demandThreshold int
}

func New(cache RedisService, stream StreamService, demandThreshold int) BeerView {
	return BeerView{cache, stream, demandThreshold}
}

func (b BeerView) Increment(ctx context.Context, beerID uint) (int64, error) {
	key := fmt.Sprintf("beer:%d:views", beerID)
	result, err := b.cache.Incr(ctx, key)
	if err != nil {
		return 0, fmt.Errorf("beerview.Increment(): %w", err)
	}

	return result, nil
}

func (b BeerView) IsHighDemandReached(ctx context.Context, beerID uint) bool {
	key := fmt.Sprintf("beer:%d:views", beerID)
	data, err := b.cache.Get(ctx, key)
	if err != nil {
		return false
	}

	views, _ := strconv.Atoi(data)

	return int64(views) >= int64(b.demandThreshold)
}

func (b BeerView) PublishHighDemand(ctx context.Context, beer model.Beer) error {
	beerId := strconv.Itoa(int(beer.ID))
	currentPrice := fmt.Sprintf("%.2f", beer.Price)
	if err := b.stream.Publish(ctx, model.NewHighDemand, []byte(beerId), []byte(currentPrice)); err != nil {
		return fmt.Errorf("beerview.PublishHighDemand(): %w", err)
	}

	return nil
}

func (b BeerView) ResetViewCounter(ctx context.Context, beerID uint) error {
	key := fmt.Sprintf("beer:%d:views", beerID)
	if err := b.cache.Set(ctx, key, 0, 0); err != nil {
		return fmt.Errorf("beerview.ResetViewCounter(): %w", err)
	}

	return nil
}
