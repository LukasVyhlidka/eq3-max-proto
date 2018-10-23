package client

import (
	"testing"

	"github.com/LukasVyhlidka/eq3-max-proto/model"

	"github.com/stretchr/testify/assert"
)

func TestClientDownloadMessages(t *testing.T) {
	messages, err := ObtainInitialMessages("192.168.95.215:62910")

	assert.Nil(t, err)
	assert.NotNil(t, messages)

	foundM := false
	for _, msg := range messages {
		if msg.GetMessageType() == "M" {
			mMsg, ok := msg.(model.MMessage)
			assert.True(t, ok)
			assert.NotNil(t, mMsg)
			foundM = true
		}
	}

	assert.True(t, foundM)
}
