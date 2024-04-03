package channels

// Mux is a multiplexer that can send and receive values of type T.
// It is used to fan-out values to multiple consumers.
// It is safe for concurrent use.
type Mux[T any] struct {
	sinks []chan T
}

// Send sends a value to all consumers.
func (m *Mux[T]) Send(t T) {
	for _, sink := range m.sinks {
		sink <- t
	}
}

// NewSink returns a new channel that can be used to receive values.
func (m *Mux[T]) NewSink() chan T {
	sink := make(chan T)
	m.sinks = append(m.sinks, sink)
	return sink
}
