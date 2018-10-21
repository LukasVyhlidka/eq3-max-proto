package parser

import (
	"encoding/binary"
	"fmt"
	"log"
	"regexp"

	"github.com/LukasVyhlidka/eq3-max-proto/model"

	b64 "encoding/base64"
)

type Error string

func (e Error) Error() string { return string(e) }

const ErrInvalidMessage Error = Error("Invalid Message")

var msgPattern *regexp.Regexp = regexp.MustCompile(`^(\w):(.+)$`)

var nilMsg = model.LMessage{}

func ParseLMessage(message string) (model.LMessage, error) {
	if !msgPattern.MatchString(message) {
		return nilMsg, ErrInvalidMessage
	}

	msgParts := msgPattern.FindStringSubmatch(message)
	msgType := msgParts[1]
	msg := msgParts[2]
	if msgType != "L" {
		return nilMsg, Error("Wrong message type: " + msgType)
	}

	data, error := b64.StdEncoding.DecodeString(msg)
	if error != nil {
		return nilMsg, ErrInvalidMessage
	}

	if len(data) < 2 {
		log.Fatal(fmt.Sprintf("Message [%s] data part is of length %d.", message, len(data)))
		return nilMsg, Error("Data part of the message has to be at least 2 bytes (0xce 0x00).")
	}

	if data[len(data)-1] != 0x00 && data[len(data)-2] != 0xce {
		log.Printf("Message [%s] data does not end with 0xce 0x00 bytes.", message)
	}

	maxDevices := []model.MaxDevice{}
	if len(data) > 2 {
		i := 0
		subMessageLength := int(data[i])

		for i < len(data) && subMessageLength > 0 {
			if (i + subMessageLength + 1) > len(data) {
				return nilMsg, Error("One SubMessage out of bounds exception.")
			}

			deviceType := model.ECO
			rfAddress := int(binary.BigEndian.Uint32([]byte{0, data[i+1], data[i+2], data[i+3]}))
			unknown := data[i+4]
			flags := int(binary.BigEndian.Uint32([]byte{0, 0, data[i+5], data[i+6]}))
			var valvePosition byte = 0
			var temperature int = 0
			var dateUntil int = 0
			var timeUntil byte = 0
			var actualTemperature int = 0

			if subMessageLength > 6 {
				deviceType = model.VALVE
				valvePosition = data[i+7]
				temperature = int(data[i+8] & 0x7F) // First bit belongs to actual temperature
				dateUntil = int(binary.BigEndian.Uint32([]byte{0, 0, data[i+9], data[i+10]}))
				timeUntil = data[i+11]
			}

			if subMessageLength > 11 {
				deviceType = model.THERMOSTAT
				actualTemperature = (((int(data[i+8]) & 0x80) << 1) | int(data[i+12]&0xFF)) // temperature first bit belongs to here
			}

			maxDevices = append(maxDevices, model.MaxDevice{
				DeviceType:        deviceType,
				RfAddress:         rfAddress,
				Unknown:           unknown,
				Flags:             flags,
				ValvePosition:     valvePosition,
				Temperature:       temperature,
				DateUntil:         dateUntil,
				TimeUntil:         timeUntil,
				ActualTemparature: actualTemperature,
			})

			i = i + 1 + subMessageLength
			if i >= len(data) {
				break
			}

			subMessageLength = int(data[i])
		}
	}

	return model.LMessage{Devices: maxDevices}, nil
}
