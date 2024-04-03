package channels

type Mux[T any] struct {
}

func (m *Mux[T]) Send(t T)
func (m *Mux[T]) NewSink() chan T
