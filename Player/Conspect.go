package Player

type Conspect struct {
	Name string
}

//func (c *Conspect) ItemMakeAction(p *Player) string {
//	return "i am kospect"
//}

func (c *Conspect) ItemAction(p *Player) string {
	return "я конспект"
}
