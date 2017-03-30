package messaging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryBus(t *testing.T) {
	b := &MemoryBus{}
	b.Start()

	err := b.Publish("topic", []byte("message1"))
	assert.NoError(t, err)
	// should be able to publish with no subscribers

	c1, err := b.Subscribe("topic", "")
	assert.NoError(t, err)

	c2, err := b.Subscribe("topic", "")
	assert.NoError(t, err)

	err = b.Publish("topic", []byte("message2"))
	assert.NoError(t, err)
	err = b.Publish("topic", []byte("message3"))
	assert.NoError(t, err)

	b.Stop()

	received := 0
	messages := []string{}
	for m := range c1 {
		received++
		messages = append(messages, string(m))
	}

	for m := range c2 {
		received++
		messages = append(messages, string(m))
	}
	assert.Equal(t, "message1", messages[0])
	assert.Equal(t, "message2", messages[4])
	assert.Equal(t, 6, len(messages))
	assert.Equal(t, 6, received)
}