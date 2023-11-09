package data

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

var DeviceOSNames = map[protocol.DeviceOS]string{
	protocol.DeviceAndroid:   "Android",
	protocol.DeviceIOS:       "iOS",
	protocol.DeviceOSX:       "OSX",
	protocol.DeviceFireOS:    "FireOS",
	protocol.DeviceGearVR:    "GearVR",
	protocol.DeviceHololens:  "Hololens",
	protocol.DeviceWin10:     "Win10",
	protocol.DeviceWin32:     "Win32",
	protocol.DeviceDedicated: "Dedicated",
	protocol.DeviceTVOS:      "TVOS",
	protocol.DeviceOrbis:     "Orbis",
	protocol.DeviceNX:        "NX",
	protocol.DeviceXBOX:      "XBOX",
	protocol.DeviceWP:        "WP",
	protocol.DeviceLinux:     "Linux",
}

// DeviceOSからOS名を取得
func GetDeviceOSName(deviceOS protocol.DeviceOS) string {
	if name, ok := DeviceOSNames[deviceOS]; ok {
		return name
	}
	return "Unknown"
}
