package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMux tests Mux type usability.
func TestMux_Send(t *testing.T) {
	m := &Mux[int]{}
	defer m.Close()

	sink1 := m.NewSink()
	sink2 := m.NewSink()

	want := 1
	go m.Send(want)
	got1 := <-sink1
	got2 := <-sink2

	assert.Equal(t, want, got1)
	assert.Equal(t, want, got2)
}

// TestMux_Close tests Mux type usability.
func TestMux_CloseSink(t *testing.T) {
	m := &Mux[int]{}
	sink1 := m.NewSink()
	sink2 := m.NewSink()

	want := 2
	got := m.Len()
	assert.Equal(t, want, got)

	m.CloseSink(sink1)
	want = 1
	got = m.Len()
	assert.Equal(t, want, got)
	<-sink1

	m.CloseSink(sink2)
	want = 0
	got = m.Len()
	assert.Equal(t, want, got)
	<-sink2
}
