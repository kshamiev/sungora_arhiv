package rabbit

type ConsumerHandler interface {
	Handler(data []byte)
}

type ConsumerHandlerFunc func(data []byte)

func (f ConsumerHandlerFunc) Handler(data []byte) {
	f(data)
}

// ////

type Consumer struct {
	Exchange string
	Tag      string
}

func NewConsumer(exchange string, tag string) *Consumer {
	return &Consumer{
		Exchange: exchange,
		Tag:      tag,
	}
}

// ////

type Producer struct {
	Exchange  string
	IsConfirm bool
}

func NewProducer(exchange string, isConfirm bool) *Producer {
	return &Producer{
		Exchange:  exchange,
		IsConfirm: isConfirm,
	}
}
