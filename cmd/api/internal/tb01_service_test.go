package internal

import (
	"circle/pkg/tb01"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	err error
}

func (m *mockRepo) Create(data *tb01.TB01) error {
	return m.err
}

func TestTB01Service_Create(t *testing.T) {
	mock := &mockRepo{}
	s := NewTB01Service(mock)

	want := fmt.Errorf("empty ColTexto field")
	err := s.Create(&tb01.TB01{})
	assert.Equal(t, err, want)

	mock.err = fmt.Errorf("failed to insert data on database")
	err = s.Create(&tb01.TB01{ColTexto: "test"})
	assert.Equal(t, err, mock.err)

	mock.err = nil
	err = s.Create(&tb01.TB01{ColTexto: "test"})
	assert.Nil(t, err)
}
