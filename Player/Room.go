package Player

type Room struct {
	INode
	RoomName    string
	RoomObjects []*RoomObject
}

// создание комнаты и добавление ее на карту
func CreateRoom(roomName string, roomObjects []*RoomObject) *Node {
	//var room = Room{
	//	RoomName:    roomName,
	//	RoomObjects: []*RoomObject{},
	//}

	var newRoom = Node{
		Name:           roomName,
		GameObjectType: Room{},
		RoomObjects:    roomObjects,
	}
	return &newRoom
}
