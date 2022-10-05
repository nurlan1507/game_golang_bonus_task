package Player

import (
	"errors"
	"fmt"
)

type Back struct {
	IItem
	Name     string
	Items    []IItem
	Capacity int
}

func (b *Back) GetItems(p *Player) string {
	fmt.Println(b.Items)
	return ""
}

func (b *Back) ItemInitAction(p *Player) string {
	p.GetBackPack(b)
	return "вы надели :рюкзак"
}

func (b *Back) GetItemName() string {
	return b.Name
}
func (b *Back) AddItem(item IItem) (*Back, error) {
	if len(b.Items) == b.Capacity {
		return nil, errors.New("inventory is full")
	}
	b.Items = append(b.Items, item)
	return b, nil
}

func (b *Back) RemoveItem(item IItem) (*Back, error) {
	var index int
	if len(b.Items) == 0 {
		return nil, errors.New("inventory is empty")
	} else {
		for i, v := range b.Items {
			if v == item {
				index = i
				b.Items[index] = b.Items[len(b.Items)-1]
				b.Items[len(b.Items)-1] = nil
				b.Items = b.Items[:len(b.Items)-1]
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
