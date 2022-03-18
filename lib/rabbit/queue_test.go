package rabbit

import (
	"context"
	"encoding/json"
	"github.com/shopspring/decimal"
	"sungora/lib/logger"
	"sungora/lib/typ"
	"testing"
)

func TestQueueProducer(t *testing.T) {
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	pub := NewProducer(context.Background(), "", true)
	err = pub.Queue("funtik", getOrders())
	if err != nil {
		t.Fatal(err)
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
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "ctxValue")

	cons := NewConsumer(ctx, "", "sample")

	err = cons.Queue("funtik", &Model{})
	if err != nil {
		t.Fatal(err)
	}

	err = cons.Cancel()
	if err != nil {
		t.Fatal(err)
	}

	CloseWait()
}

type Model struct{}

func (m *Model) RBCHandler(ctx context.Context, data []byte) {
	if key, ok := ctx.Value("key").(string); !ok || key != "ctxValue" {
		panic("context no delivery")
	} else {
		logger.Gist(ctx).Infoln(key)
		o := Order{}
		_ = json.Unmarshal(data, &o)
		logger.Gist(ctx).Infoln(o)
	}
}

type Order struct {
	ID    typ.UUID        `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
}

func getOrders() []interface{} {
	return []interface{}{
		Order{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 1",
			Price: decimal.NewFromInt(23),
		},
		Order{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 2",
			Price: decimal.NewFromInt(657),
		},
		Order{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 3",
			Price: decimal.NewFromInt(856),
		},
		Order{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 4",
			Price: decimal.NewFromInt(387),
		},
		Order{
			ID:    typ.UUIDNew(),
			Name:  "Popcorn 5",
			Price: decimal.NewFromInt(884),
		},
	}
}
