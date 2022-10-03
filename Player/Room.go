package Player

type Room struct {
	INode
	RoomName string
}

// создание комнаты и добавление ее на карту
func CreateRoom(roomName string) *Node {
	//var room = Room{
	//	RoomName:    roomName,
	//	RoomObjects: []*RoomObject{},
	//}

	var newRoom = Node{
		Name:           roomName,
		GameObjectType: Room{},
	}
	return &newRoom
}
