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

// Len returns the number of consumers.
func (m *Mux[T]) Len() int {
	return len(m.sinks)
}

// Close closes all consumers.
func (m *Mux[T]) Close() {
	for _, sink := range m.sinks {
		close(sink)
	}
	m.sinks = nil
}

// CloseSink closes a specific consumer.
func (m *Mux[T]) CloseSink(sink chan T) {
	for i, s := range m.sinks {
		if s == sink {
			close(s)
			m.sinks = append(m.sinks[:i], m.sinks[i+1:]...)
			return
		}
	}
}
