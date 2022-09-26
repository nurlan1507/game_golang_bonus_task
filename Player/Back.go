package Player

import (
	"errors"
	"fmt"
	"strings"
)

type Back struct {
	Item
	Items    []string
	Capacity int
}

func (b *Back) ItemMakeAction(p *Player) string {
	var newBack = Back{
		Items:    []string{},
		Capacity: 10,
	}
	var newInventory = new(Inventory)
	p.Inventory = newInventory
	fmt.Println("был создан инвентарь размером" + string(rune(newBack.Capacity)))
	return ""

}

func (b *Back) AddItem(item string) (*Back, error) {
	if len(b.Items) == b.Capacity {
		return nil, errors.New("inventory is full")
	}
	b.Items = append(b.Items, item)
	return b, nil
}

func (b *Back) RemoveItem(item string) (*Back, error) {
	if len(b.Items) == 0 {
		return nil, errors.New("inventory is empty")
	} else {
		for i, v := range b.Items {
			if v == item {
				b.Items = append(b.Items[:i], b.Items[i+1])
				return b, nil
			}
		}
	}
	return nil, errors.New("some error occured")
}

func (b *Back) PrintAll() string {
	return strings.Join(b.Items, ",")
}
