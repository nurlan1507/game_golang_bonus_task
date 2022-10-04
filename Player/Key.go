package Player

// key
type Key struct {
	Item
	Name string
}

func (k *Key) ItemAction(p *Player) string {
	return ""
}
