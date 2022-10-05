package Player

type IItem interface {
	ItemInitAction(p *Player) string
	ItemAction(p *Player) string
	GetItems(p *Player) string
	GetItemName() string
}
