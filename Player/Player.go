package Player

import (
	"fmt"
)

type IInventory interface {
	AddItem(item string) (*Back, error)
	DeleteItem(item string) (*Back, error)
	PrintItems() string
}
type IPlayer interface {
	ChangeLocation()
}

type Inventory struct {
	IInventory
	Items *Back
}

func (I *Inventory) AddItem(item string) (*Back, error) {
	back, err := I.Items.AddItem(item)
	if err != nil {
		return nil, err
	}
	return back, nil
}

func (I *Inventory) DeleteItem(item string) (*Back, error) {
	back, err := I.Items.RemoveItem(item)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("\n вещь %v была добавлена в инвентарь \n", item)
	return back, nil
}

func (I *Inventory) PrintItems() string {
	if I != nil {
		return I.Items.PrintAll()
	}
	return "no backpack"
}

type Player struct {
	nickname  string
	Inventory *Inventory
}

func NewPlayer(nickname string) *Player {
	var newPlayer = Player{
		nickname:  nickname,
		Inventory: nil,
	}
	return &newPlayer
}
func (P *Player) GetBackPack() {
	var newBack = Back{
		items:    []string{},
		capacity: 10,
	}
	var newInventory = new(Inventory)
	newInventory.Items = &newBack
	P.Inventory = newInventory
	fmt.Println("был создан инвентарь размером")
}

func (P *Player) ChangeLocation() {

}

func (P *Player) GetPlayerInfo() {
	fmt.Printf("\n player name: %+v", P.nickname)
	fmt.Printf("\n inventory: %v \n", P.Inventory.PrintItems())
}
