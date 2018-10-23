package model

type Message interface {
	GetMessageType() string
}

type GenericMessage struct {
	Message     string
	MessageType string
}

func (t GenericMessage) GetMessageType() string {
	return t.MessageType
}
