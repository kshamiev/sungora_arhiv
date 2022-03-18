package rabbit

import (
	"context"
	"log"
	"testing"
)

func TestQueueProducer(t *testing.T) {
	err := Init(&Config{
		Uri: "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		t.Fatal(err)
	}

	data := []string{
		"one",
		"two",
		"three",
		"popcorn",
		"Tutorials",
		"RabbitMQ",
	}

	pub := NewProducer(context.Background(), "", true)
	err = pub.Queue("funtik", data...)
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
	ctx = context.WithValue(ctx, "key", "value")

	var h ConsumerHandlerFunc = func(ctx context.Context, data []byte) {
		if "valgue" != ctx.Value("key").(string) {
			//panic("context no delivery")
			log.Fatal("context no delivery")
		}
		log.Println(string(data))
	}

	cons := NewConsumer(ctx, "", "sample")
	err = cons.Queue("funtik", h)
	if err != nil {
		t.Fatal(err)
	}
	err = cons.Cancel()
	if err != nil {
		t.Fatal(err)
	}

	CloseWait()
}
