// alias direct (worker pull, transaction)
package rabbit

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func TestQueueProducer(t *testing.T) {
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	pub, err := NewProducerQueue("test_queue")
	if err != nil {
		t.Fatal(err)
	}

	for _, o := range getOrders() {
		err = pub.Queue(o)
		if err != nil {
			t.Fatal(err)
		}
	}

	CloseWait()
}

func TestQueueConsumer(t *testing.T) {
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	cons, err := NewConsumerQueue("test_queue")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "ctxValue")
	var h = func(ctx context.Context, data []byte) {
		if key, ok := ctx.Value("key").(string); !ok || key != "ctxValue" {
			panic("context no delivery")
		} else {
			t.Log(key, string(data))
		}
	}

	err = cons.Handler(ctx, h)
	if err != nil {
		t.Fatal(err)
	}

	err = cons.Cancel()
	if err != nil {
		t.Fatal(err)
	}

	CloseWait()
}

type Order struct {
	ID    uuid.UUID       `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
}

func getOrders() []*Order {
	return []*Order{
		{
			ID:    uuid.New(),
			Name:  "Popcorn 1",
			Price: decimal.NewFromInt(23),
		},
		{
			ID:    uuid.New(),
			Name:  "Popcorn 2",
			Price: decimal.NewFromInt(657),
		},
		{
			ID:    uuid.New(),
			Name:  "Popcorn 3",
			Price: decimal.NewFromInt(856),
		},
		{
			ID:    uuid.New(),
			Name:  "Popcorn 4",
			Price: decimal.NewFromInt(387),
		},
		{
			ID:    uuid.New(),
			Name:  "Popcorn 5",
			Price: decimal.NewFromInt(884),
		},
		{
			ID:    uuid.New(),
			Name:  "Popcorn 6",
			Price: decimal.NewFromInt(165),
		},
		{
			ID:    uuid.New(),
			Name:  "Popcorn 7",
			Price: decimal.NewFromInt(44),
		},
	}
}
