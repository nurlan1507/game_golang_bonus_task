package Player

type Conspect struct {
	Item
	Name string
}

func (c *Conspect) ItemMakeAction(p *Player) string {
	return "i am kospect"
}
