package Locations

type Room struct {
	INode
	RoomName    string
	RoomObjects []*RoomObject
}

// создание комнаты и добавление ее на карту
func CreateRoom(roomName string) *Node {
	var room = Room{
		RoomName:    roomName,
		RoomObjects: []*RoomObject{},
	}
	var newRoom = Node{
		GameObject: room,
		name:       roomName,
	}
	return &newRoom
}
