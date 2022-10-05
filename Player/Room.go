package Player

type Room struct {
	INode
	RoomName string
}

// создание комнаты и добавление ее на карту
func CreateRoom(roomName string) *Node {

	var newRoom = Node{
		Name:           roomName,
		GameObjectType: Room{},
		Door:           nil,
		RoomObjects:    map[string][]IItem{},
	}
	return &newRoom
}
