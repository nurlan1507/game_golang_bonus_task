package Player

import (
	"errors"
)

type Back struct {
	Item
	Items    []*Item
	Capacity int
}

func (b *Back) ItemInitAction(p *Player) string {
	p.GetBackPack(b)
	return "ad"
}

func (b *Back) AddItem(item *Item) (*Back, error) {
	if len(b.Items) == b.Capacity {
		return nil, errors.New("inventory is full")
	}
	b.Items = append(b.Items, item)
	return b, nil
}

func (b *Back) RemoveItem(item *Item) (*Back, error) {
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

//
//func (b *Back) PrintAll() string {
//	return strings.Join(b.Items, ",")
//}
