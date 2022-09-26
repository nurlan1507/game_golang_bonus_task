package Player

type RoomObject struct {
	INode
	RoomObjectName string
	Items          []*Item
}

func CreateRoomObject(roomObjectName string) *RoomObject {
	var newRoomObject = RoomObject{
		RoomObjectName: roomObjectName,
		Items:          []*Item{},
	}
	return &newRoomObject
}
