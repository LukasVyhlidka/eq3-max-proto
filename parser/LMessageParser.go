package parser

import (
	"github.com/LukasVyhlidka/eq3-max-proto/model"
)

type Error string

func (e Error) Error() string { return string(e) }

const ErrInvalidMessage Error = Error("Invalid Message")

func ParseLMessage(message string) (model.LMessage, error) {
	return model.LMessage{}, Error("Not implemented, yet.")
}
