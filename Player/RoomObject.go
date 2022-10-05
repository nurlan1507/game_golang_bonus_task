package Player

type RoomObject struct {
	INode
	RoomObjectName string
	Items          []*IItem
}

func CreateRoomObject(roomObjectName string) *RoomObject {
	var newRoomObject = RoomObject{
		RoomObjectName: roomObjectName,
		Items:          []*IItem{},
	}
	return &newRoomObject
}
