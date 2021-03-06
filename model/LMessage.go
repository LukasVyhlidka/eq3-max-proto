package model

import "fmt"

type LMessage struct {
	// Devices - slice of Max Devices.
	Devices []MaxDevice
}

type MaxDevice struct {
	DeviceType        MaxDeviceType
	RfAddress         int
	Unknown           byte
	Flags             int
	ValvePosition     byte
	Temperature       int
	DateUntil         int
	TimeUntil         byte
	ActualTemparature int
}

// MaxDeviceType = Type of the device
type MaxDeviceType int

const (
	// ECO = Button device type
	ECO MaxDeviceType = 0
	// VALVE = Radiator Valve
	VALVE MaxDeviceType = 1
	// THERMOSTAT = Wall Thermostat
	THERMOSTAT MaxDeviceType = 2
)

func (t MaxDevice) String() string {
	return fmt.Sprintf(
		"MaxDevice{DeviceType: %d, RfAddress: %d, Unknown: %d, Flags: %d, ValvePosition: %d, Temperature: %d, DateUntil: %d, TimeUntil: %d ActualTemperature: %d}",
		t.DeviceType, t.RfAddress, t.Unknown, t.Flags, t.ValvePosition, t.Temperature, t.DateUntil, t.TimeUntil, t.ActualTemparature)
}

func (m LMessage) GetMessageType() string {
	return "L"
}
