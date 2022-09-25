package Player

import "game/nurlan/Locations"

type Key struct {
	Locations.Item
}

func (k *Key) ItemMakeAction() string {
	return "дверь открыта"
}
