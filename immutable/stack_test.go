package immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Int(t *testing.T) {
	var s Stack
	assert.True(t, s.IsEmpty())
	assert.Nil(t, s.Peek())
	s1 := s.Push(1)
	assert.Nil(t, s.Peek())
	assert.False(t, s1.IsEmpty())
	assert.Equal(t, 1, s1.Peek())
}

func TestStack_String(t *testing.T) {
	var s Stack
	assert.True(t, s.IsEmpty())
	s2 := s.Push("foo")
	assert.True(t, s.IsEmpty())
	assert.False(t, s2.IsEmpty())
	assert.Equal(t, s2.Peek(), "foo")
	s3 := s2.Push("bar")
	assert.Equal(t, s3.Peek(), "bar")
	assert.Equal(t, s3.Pop().Peek(), "foo")
	assert.Equal(t, s3.Peek(), "bar")
	s4 := s3.Pop().Push("bar")
	assert.False(t, s3.IsEmpty())
	assert.Equal(t, s4.Peek(), "bar")
	s5 := s4.Push("foo").Push("fee")
	assert.Equal(t, s5.Pop().Pop().Peek(), "bar")
}
