package rbr

type Exchange string

const (
	Exchange1 Exchange = "Exchange1"
	Exchange2 Exchange = "Exchange2"
	Exchange3 Exchange = "Exchange3"
)

func (e Exchange) String() string {
	return string(e)
}
