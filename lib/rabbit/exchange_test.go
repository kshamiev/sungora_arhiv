package rabbit

import (
	"context"
	"testing"
)

func TestExchangeProducer(t *testing.T) {
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	data := []interface{}{
		"one",
		"two",
		"three",
		"popcorn",
		"Tutorials",
		"RabbitMQ",
	}

	pub := NewProducer(context.Background(), "test-exchange", true)
	err = pub.Exchange("test-key-one", "test-queue-one", data)
	if err != nil {
		t.Fatal(err)
	}

	CloseWait()
}

func TestExchangeConsumer(t *testing.T) {
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "ctxValue")

	var h ConsumerHandlerFunc = func(ctx context.Context, data []byte) {
		if key, ok := ctx.Value("key").(string); !ok || key != "ctxValue" {
			panic("context no delivery")
		} else {
			t.Log(key, string(data))
		}
	}

	cons := NewConsumer(ctx, "test-exchange", "sample")
	err = cons.Exchange("test-key-one", "test-queue-one", h)
	if err != nil {
		t.Fatal(err)
	}
	err = cons.Cancel()
	if err != nil {
		t.Fatal(err)
	}

	CloseWait()
}

//func Handle(ctx context.Context, data []byte) {
//	if key, ok := ctx.Value("key").(string); !ok || key != "value" {
//		panic("context no delivery")
//	} else {
//		logger.Gist(ctx).Infoln(key, string(data))
//	}
//}
