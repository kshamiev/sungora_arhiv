// alias topic
package rabbit

import (
	"context"
	"testing"
)

func TestTopicProducer(t *testing.T) {
	t.Skip()
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	pub, err := NewProducerTopic("test_exchange_topic")
	if err != nil {
		t.Fatal(err)
	}

	for _, o := range getOrders() {
		err = pub.Topic("one.two.three.four", o)
		if err != nil {
			t.Fatal(err)
		}
	}

	CloseWait()
}

func TestTopicConsumer(t *testing.T) {
	t.Skip()
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	cons, err := NewConsumerTopic("test_exchange_topic", "test_queue_topic", []string{"*.two.#", "#.four"})
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
