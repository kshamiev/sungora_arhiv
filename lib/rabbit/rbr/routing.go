package rbr

type RoutingKey string

const (
	Key1 RoutingKey = "Key1"
	Key2 RoutingKey = "Key2"
	Key3 RoutingKey = "Key3"
)

func (r RoutingKey) String() string {
	return string(r)
}
