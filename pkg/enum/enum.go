package enum

import "fmt"

type cameraStatus int

const (
	cameraOffline cameraStatus = iota
	cameraOnline
	CameraNotResponding
)

// using map
var uMap = map[cameraStatus]string{
	cameraOffline:       "offline",
	cameraOnline:        "online",
	CameraNotResponding: "NotResponding",
}

func (cs cameraStatus) String() string {
	switch cs {
	case cameraOffline:
		return "offline"
	case cameraOnline:
		return "online"
	case CameraNotResponding:
		return "Not Responding"
	default:
		return "Internal Server Error"
	}
}

func Enum() {
	var invalidStatus cameraStatus = 10
	fmt.Println(uMap[cameraOffline])
	fmt.Println(cameraOnline)
	fmt.Println(invalidStatus)
}
