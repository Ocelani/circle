package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMux tests Mux type usability.
func TestMux_Send(t *testing.T) {
	m := &Mux[int]{}
	sink1 := m.NewSink()
	sink2 := m.NewSink()

	want := 1
	go m.Send(want)

	got1 := <-sink1
	got2 := <-sink2

	assert.Equal(t, want, got1)
	assert.Equal(t, want, got2)
}
