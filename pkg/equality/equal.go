package equality

type Equalizer interface {
	Equal(object interface{}) bool
}
