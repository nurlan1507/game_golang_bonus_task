package Player

import (
	"errors"
	"strings"
)

type Back struct {
	items    []string
	capacity int
}

//func ItemMakeAction(p *Player) *Back {
//	var newBack = Back{
//		items:    []string{},
//		capacity: 10,
//	}
//	var newInventory = new(Inventory)
//	p.Inventory = newInventory
//	fmt.Println("был создан инвентарь размером" + string(newBack.capacity))
//	return &newBack
//}

func (b *Back) AddItem(item string) (*Back, error) {
	if len(b.items) == b.capacity {
		return nil, errors.New("inventory is full")
	}
	b.items = append(b.items, item)
	return b, nil
}

func (b *Back) RemoveItem(item string) (*Back, error) {
	if len(b.items) == 0 {
		return nil, errors.New("inventory is empty")
	} else {
		for i, v := range b.items {
			if v == item {
				b.items = append(b.items[:i], b.items[i+1])
				return b, nil
			}
		}
	}
	return nil, errors.New("some error occured")
}

func (b *Back) PrintAll() string {
	return strings.Join(b.items, ",")
}
