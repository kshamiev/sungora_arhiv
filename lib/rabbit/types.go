package rabbit

import (
	"context"
	"sync"
)

type ConsumerHandler interface {
	RBCHandler(ctx context.Context, data []byte)
}

type ConsumerHandlerFunc func(ctx context.Context, data []byte)

func (f ConsumerHandlerFunc) RBCHandler(ctx context.Context, data []byte) {
	f(ctx, data)
}

// ////

type Consumer struct {
	exchange string
	tag      string
	ctx      context.Context
	cnt      int
	wg       sync.WaitGroup
}

func NewConsumer(ctx context.Context, exchange string, tag string) *Consumer {
	return &Consumer{
		exchange: exchange,
		tag:      tag,
		ctx:      ctx,
	}
}

// ////

type Producer struct {
	exchange  string
	isConfirm bool
	ctx       context.Context
}

func NewProducer(ctx context.Context, exchange string, isConfirm bool) *Producer {
	return &Producer{
		exchange:  exchange,
		isConfirm: isConfirm,
		ctx:       ctx,
	}
}
