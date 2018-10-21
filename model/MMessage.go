package model

type MMessage struct {
	Index   int
	Count   int
	Rooms   []MaxRoomMeta
	Devices []MaxDeviceMeta
}

type MaxRoomMeta struct {
	ID        int
	Name      string
	RfAddress int
}

type MaxDeviceMeta struct {
	Type         DeviceTypeMeta
	RfAddress    int
	SerialNumber string
	Name         string
	RoomID       int
}

// MaxDeviceType = Type of the device
type DeviceTypeMeta int

const (
	Cube                  DeviceTypeMeta = 0
	HeatingThermostat     DeviceTypeMeta = 1
	HeatingThermostatPlus DeviceTypeMeta = 2
	WallMountedThermostat DeviceTypeMeta = 3
	ShutterContact        DeviceTypeMeta = 4
	PushButton            DeviceTypeMeta = 5
)
