// alias direct (worker pull, transaction)
package rabbit

import (
	"context"
	"sungora/lib/typ"
	"testing"

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

	err = cons.Queue(ctx, h)
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
	ID    typ.UUID        `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
}

func getOrders() []*Order {
	return []*Order{
		{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 1",
			Price: decimal.NewFromInt(23),
		},
		{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 2",
			Price: decimal.NewFromInt(657),
		},
		{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 3",
			Price: decimal.NewFromInt(856),
		},
		{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 4",
			Price: decimal.NewFromInt(387),
		},
		{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 5",
			Price: decimal.NewFromInt(884),
		},
		{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 6",
			Price: decimal.NewFromInt(165),
		},
		{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 7",
			Price: decimal.NewFromInt(44),
		},
	}
}
