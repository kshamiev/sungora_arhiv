package rabbit

import (
	"sungora/lib/rabbit/rbr"
	"testing"
)

func TestQueueProducer(t *testing.T) {
	err := Init(&Config{
		Uri:      "amqp://guest:guest@localhost:5672/",
		Exchange: rbr.Exchange1.String(),
	})
	if err != nil {
		t.Fatal(err)
	}

	data := []string{
		"one",
		"two",
		"three",
	}

	pub := NewProducer("", true)
	err = pub.Queue("funtik", data...)
	if err != nil {
		t.Fatal(err)
	}

	CloseWait()
}

func TestQueueConsumer(t *testing.T) {
	err := Init(&Config{
		Uri:      "amqp://guest:guest@localhost:5672/",
		Exchange: rbr.Exchange1.String(),
	})
	if err != nil {
		t.Fatal(err)
	}

	var h ConsumerHandlerFunc = func(data []byte) {
		t.Log(string(data))
	}

	cons := NewConsumer("", "sample2221")
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
