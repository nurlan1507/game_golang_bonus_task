package Player

type Door struct {
	Item
	Status bool
	Name   string
}

func GenerateDoor(to *Node) *Door {
	newDoor := Door{Status: false, Name: "дверь"}
	return &newDoor
}
