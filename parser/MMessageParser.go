package parser

import (
	b64 "encoding/base64"
	"encoding/binary"
	"regexp"
	"strconv"

	"github.com/LukasVyhlidka/eq3-max-proto/model"
)

var mMsgPattern *regexp.Regexp = regexp.MustCompile(`^(\d{2}),(\d{2}),([^,]+)$`)

var nilMMsg = model.MMessage{}

func ParseMMessage(message string) (model.MMessage, error) {
	if !msgPattern.MatchString(message) {
		return nilMMsg, ErrInvalidMessage
	}

	msgParts := msgPattern.FindStringSubmatch(message)
	msgType := msgParts[1]
	msg := msgParts[2]

	if msgType != "M" {
		return nilMMsg, Error("Wrong message type: " + msgType)
	}

	dataParts := mMsgPattern.FindStringSubmatch(msg)
	if dataParts == nil {
		return nilMMsg, ErrInvalidMessage
	}

	index, indexErr := strconv.ParseInt(dataParts[1], 16, 16)
	count, countErr := strconv.ParseInt(dataParts[2], 16, 16)
	data, dataError := b64.StdEncoding.DecodeString(dataParts[3])
	var dataIndex int = 2

	if indexErr != nil || countErr != nil || dataError != nil {
		return nilMMsg, Error("Unexpected error when parsing the data part.")
	}

	// Rooms
	rooms := []model.MaxRoomMeta{}
	roomCount := int(data[dataIndex] & 0xff)
	dataIndex++
	for i := 0; i < roomCount; i++ {
		roomID := int(data[dataIndex] & 0xff)
		dataIndex++
		roomNameLength := int(data[dataIndex] & 0xff)
		dataIndex++
		roomName := string(data[dataIndex : dataIndex+roomNameLength])
		dataIndex += roomNameLength
		rfAddress := int(binary.BigEndian.Uint32([]byte{0, data[dataIndex], data[dataIndex+1], data[dataIndex+2]}))
		dataIndex += 3

		rooms = append(rooms, model.MaxRoomMeta{
			ID:        roomID,
			Name:      roomName,
			RfAddress: rfAddress,
		})
	}

	// Devices
	devices := []model.MaxDeviceMeta{}
	deviceCount := int(data[dataIndex] & 0xff)
	dataIndex++
	for i := 0; i < deviceCount; i++ {
		var deviceType model.DeviceTypeMeta = model.DeviceTypeMeta(int(data[dataIndex] & 0xff))
		dataIndex++
		rfAddress := int(binary.BigEndian.Uint32([]byte{0, data[dataIndex], data[dataIndex+1], data[dataIndex+2]}))
		dataIndex += 3
		serialNo := string(data[dataIndex : dataIndex+10])
		dataIndex += 10

		deviceNameLength := int(data[dataIndex] & 0xff)
		dataIndex++
		deviceName := string(data[dataIndex : dataIndex+deviceNameLength])
		dataIndex += deviceNameLength

		roomId := int(data[dataIndex] & 0xff)
		dataIndex++

		devices = append(devices, model.MaxDeviceMeta{
			Type:         deviceType,
			RfAddress:    rfAddress,
			SerialNumber: serialNo,
			Name:         deviceName,
			RoomID:       roomId,
		})
	}

	return model.MMessage{
		Count:   int(count),
		Index:   int(index),
		Rooms:   rooms,
		Devices: devices,
	}, nil

}
