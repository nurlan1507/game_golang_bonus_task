package Player

// key
type Key struct {
	Item
	name string
}

func (k *Key) ItemMakeAction(p *Player) string {
	return "i am a key"
}
