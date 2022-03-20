// alias fanout (subscribe, logs)
package rabbit

import (
	"context"
	"testing"
)

func TestFanout(t *testing.T) {
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Consumer
	cons, err := NewConsumerTopic("test_exchange_logs", "", []string{"#"})
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

	// Producer
	pub, err := NewProducerTopic("test_exchange_logs")
	if err != nil {
		t.Fatal(err)
	}

	for _, o := range getOrders() {
		err = pub.Topic("one.two.three.four", o)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Consumer
	err = cons.Cancel()
	if err != nil {
		t.Fatal(err)
	}

	CloseWait()
}
