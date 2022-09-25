package Locations

type INode interface {
	ItemMakeAction() string
}

type Room struct {
	INode
	RoomName    string
	RoomObjects []*RoomObject
}

type RoomObject struct {
	INode
	RoomObjectName string
	Items          []*Item
}
type Item struct {
	INode
	itemName string
}

func (item *Item) ItemMakeAction() string {
	return "some item"
}
