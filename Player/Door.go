package Player

type Door struct {
	Status bool
	Name   string
}

func GenerateDoor(to *Node) *Door {
	newDoor := Door{Status: false, Name: "дверь"}
	return &newDoor
}
