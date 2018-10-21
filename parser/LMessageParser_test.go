package parser

import (
	"testing"

	"github.com/LukasVyhlidka/eq3-max-proto/model"

	"github.com/stretchr/testify/assert"
)

func TestParseEmptyString(t *testing.T) {
	_, err := ParseLMessage("")

	assert.Equal(t, ErrInvalidMessage, err)
}

func TestBadFormat(t *testing.T) {
	_, err := ParseLMessage("Hello")

	assert.Equal(t, ErrInvalidMessage, err)
}

func TestBlank(t *testing.T) {
	_, err := ParseLMessage(" ")

	assert.Equal(t, ErrInvalidMessage, err)
}

func TestDifferentMessageType(t *testing.T) {
	_, err := ParseLMessage("C:blahblahblah")

	assert.Equal(t, Error("Wrong message type: C"), err)
}

func TestExample1(t *testing.T) {
	msg, err := ParseLMessage("L:Cw/a7QkSGBgoAMwACw/DcwkSGBgoAM8ACw/DgAkSGBgoAM4A")

	assert.NoError(t, err)

	assert.NotNil(t, msg.Devices)
	assert.Equal(t, 3, len(msg.Devices))

	device1 := msg.Devices[0]
	device2 := msg.Devices[1]
	device3 := msg.Devices[2]

	assert.Equal(t, model.VALVE, device1.DeviceType)
	assert.Equal(t, 0x0FDAED, device1.RfAddress)

	assert.Equal(t, model.VALVE, device2.DeviceType)
	assert.Equal(t, 0x0fc373, device2.RfAddress)

	assert.Equal(t, model.VALVE, device3.DeviceType)
	assert.Equal(t, 0x0fc380, device3.RfAddress)
}

func TestHomeExample(t *testing.T) {
	msg, err := ParseLMessage("L:DBaBdQkSGAQmAAAA7gsVFl7xEhgAJgAAAAsVFXrxEhgAJgAAAAwWetakEhgELgAAAOcLFRZW7BIYAS4A5wALFRYj7BIYAC4AAAA=")

	assert.Nil(t, err)
	assert.NotNil(t, msg.Devices)
}

func TestRealMessageWitouth00AtEnd(t *testing.T) {
	msg, err := ParseLMessage("L:DBaBdQkSGAQuAAAA1wsVFlYJEhhkLQAAAAsVFLAJEhgHKgDTAAwWgQYJEhgEKgAAANMLFRZeCRIYZC4AAAALFRYjCRIYZC0AAAAMFnrWCRIYBC0AAADM")

	assert.NoError(t, err)
	assert.NotNil(t, msg.Devices)

	assert.Equal(t, 7, len(msg.Devices))

	assert.Equal(t, model.MaxDevice{
		DeviceType:        model.THERMOSTAT,
		RfAddress:         1474933,
		Unknown:           9,
		Flags:             4632,
		ValvePosition:     4,
		Temperature:       46,
		DateUntil:         0,
		TimeUntil:         0,
		ActualTemparature: 215,
	}, msg.Devices[0])

}

type MyStruct struct {
	f1    int
	f2    string
	inner *Inner
}

type Inner struct {
	num int
}

func TestStructEquality(t *testing.T) {
	s1 := MyStruct{f1: 42, f2: "fortytwo", inner: &Inner{num: 1}}
	s2 := MyStruct{f1: 42, f2: "fortytwo", inner: &Inner{num: 1}}

	assert.Equal(t, s1, s2)
}
