package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxMessageString(t *testing.T) {
	var d1 = MaxDevice{
		DeviceType:        VALVE,
		RfAddress:         42,
		Unknown:           1,
		Flags:             3,
		ValvePosition:     8,
		Temperature:       24,
		DateUntil:         128,
		ActualTemparature: 22,
	}

	var str = d1.String()

	assert.Equal(
		t,
		"MaxDevice{DeviceType: 1, RfAddress: 42, Unknown: 1, Flags: 3, ValvePosition: 8, Temperature: 24, DateUntil: 128, ActualTemperature: 22}",
		str)
}
