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

func (s Stream) Publish(ctx context.Context, event model.Event, body []byte) error {
	if err := s.client.XAdd(ctx, &redis.XAddArgs{
		Stream: s.streamName,
		Values: map[model.Event]interface{}{"event": event, "beer_id": body},
	}).Err(); err != nil {
		return err
	}

	log.Printf("Event %s published", event)

	return nil

}
