package Player

type INode interface {
	ItemMakeAction(p *Player) string
	AddObject()
	RemoveObject()
}
