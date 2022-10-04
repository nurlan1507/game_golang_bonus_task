package Player

type IItem interface {
	ItemInitAction(p *Player) string
	ItemAction(p *Player) string
}

type Item struct {
	IItem
	ItemName     string
	ItemInstance IItem
}

func CreateItem(itemName string, itemType *Back) *Item {
	newItem := Item{ItemInstance: itemType, ItemName: itemName}
	return &newItem
}
