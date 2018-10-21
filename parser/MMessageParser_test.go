package parser

import (
	"testing"

	"github.com/LukasVyhlidka/eq3-max-proto/model"

	"github.com/stretchr/testify/assert"
)

func TestMParseEmptyString(t *testing.T) {
	_, err := ParseMMessage("")

	assert.Equal(t, ErrInvalidMessage, err)
}

func TestMBadFormat(t *testing.T) {
	_, err := ParseMMessage("Hello")

	assert.Equal(t, ErrInvalidMessage, err)
}

func TestMBlank(t *testing.T) {
	_, err := ParseMMessage(" ")

	assert.Equal(t, ErrInvalidMessage, err)
}

func TesMtDifferentMessageType(t *testing.T) {
	_, err := ParseMMessage("C:blahblahblah")

	assert.Equal(t, Error("Wrong message type: C"), err)
}

func TestParse(t *testing.T) {
	msg, err := ParseMMessage("M:00,01,VgIDAQtMaXZpbmcgUm9vbRUWXgIIQmF0aHJvb20VFlYDB0JlZHJvb20VFLAIAxaBdU5FUTEyMDM3MTkYTGl2aW5nIFJvb20gVGhlcm1vc3RhdCAxAQEVFl5NRVExODE4MjEyGExpdmluZyBSb29tIFJhZGlhdG9yIEJpZwEBFRV6TUVRMTgxNzk4MhpMaXZpbmcgUm9vbSBSYWRpYXRvciBTbWFsbAEDFnrWTkVRMTIwMTU5MhhCYXRocm9vbSBXYWxsIFRoZXJtb3N0YXQCARUWVk1FUTE4MTgyMTAXQmF0aHJvb20gUmFkaWF0b3IgTGFkZXICARUWI01FUTE4MTgxNTAcQmF0aHJvb20gUmFkaWF0b3IgVGhlcm1vc3RhdAIDFoEGTkVRMTIwMzE2ORdXYWxsIFRoZXJtb3N0YXQgQmVkcm9vbQMBFRSwTUVRMTgxNzgyMRtSYWRpYXRvciBUaGVybW9zdGF0IEJlZHJvb20DAQ==")
	assert.NoError(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, model.MMessage{
		Count: 1,
		Index: 0,
		Rooms: []model.MaxRoomMeta{
			model.MaxRoomMeta{ID: 1, Name: "Living Room", RfAddress: 1381982},
			model.MaxRoomMeta{ID: 2, Name: "Bathroom", RfAddress: 1381974},
			model.MaxRoomMeta{ID: 3, Name: "Bedroom", RfAddress: 1381552},
		},
		Devices: []model.MaxDeviceMeta{
			model.MaxDeviceMeta{Type: model.WallMountedThermostat, RfAddress: 1474933, SerialNumber: "NEQ1203719", Name: "Living Room Thermostat 1", RoomID: 1},
			model.MaxDeviceMeta{Type: model.HeatingThermostat, RfAddress: 1381982, SerialNumber: "MEQ1818212", Name: "Living Room Radiator Big", RoomID: 1},
			model.MaxDeviceMeta{Type: model.HeatingThermostat, RfAddress: 1381754, SerialNumber: "MEQ1817982", Name: "Living Room Radiator Small", RoomID: 1},
			model.MaxDeviceMeta{Type: model.WallMountedThermostat, RfAddress: 1473238, SerialNumber: "NEQ1201592", Name: "Bathroom Wall Thermostat", RoomID: 2},
			model.MaxDeviceMeta{Type: model.HeatingThermostat, RfAddress: 1381974, SerialNumber: "MEQ1818210", Name: "Bathroom Radiator Lader", RoomID: 2},
			model.MaxDeviceMeta{Type: model.HeatingThermostat, RfAddress: 1381923, SerialNumber: "MEQ1818150", Name: "Bathroom Radiator Thermostat", RoomID: 2},
			model.MaxDeviceMeta{Type: model.WallMountedThermostat, RfAddress: 1474822, SerialNumber: "NEQ1203169", Name: "Wall Thermostat Bedroom", RoomID: 3},
			model.MaxDeviceMeta{Type: model.HeatingThermostat, RfAddress: 1381552, SerialNumber: "MEQ1817821", Name: "Radiator Thermostat Bedroom", RoomID: 3},
		},
	}, msg)
}
