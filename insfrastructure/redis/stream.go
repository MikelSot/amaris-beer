package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"

	"github.com/MikelSot/amaris-beer/model"
)

type Stream struct {
	client     *redis.Client
	streamName string
}

func NewStream(client *redis.Client, streamName string) Stream {
	return Stream{
		client:     client,
		streamName: streamName,
	}
}

func (s Stream) Publish(ctx context.Context, event model.Event, beerId []byte, currentPrice []byte) error {
	if err := s.client.XAdd(ctx, &redis.XAddArgs{
		Stream: s.streamName,
		Values: map[string]interface{}{"event": string(event), "beer_id": beerId, "current_price": currentPrice},
	}).Err(); err != nil {
		return err
	}

	log.Printf("Event %s published", event)

	return nil

}
