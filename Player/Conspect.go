package Player

type Conspect struct {
	IItem
	Name string
}

func (k *Conspect) GetItemName() string {
	return k.Name
}
func (k *Conspect) ItemInitAction(p *Player) string {
	_, err := p.Inventory.AddItem(k)
	//fmt.Println(I.ItemName)
	if err != nil {
		return err.Error()
	}
	return "добавлено в инвентарь: " + k.Name
}

// конкретное действия предмета (ключ-открыть и тд)
func (k *Conspect) ItemAction(p *Player) string {
	if p.Location.Door == nil {
		return "не к чему применить"
	}
	p.Location.Door.Status = true // true - открыта
	_, err := p.Inventory.RemoveItem(k)
	if err != nil {
		return "невозможно удалить"
	}
	return "дверь открыта"
}
