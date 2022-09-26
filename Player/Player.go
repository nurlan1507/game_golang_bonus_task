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
	Nickname  string
	Inventory *Inventory
	Location  *Node
}

func NewPlayer(nickname string) *Player {
	var newPlayer = Player{
		Nickname:  nickname,
		Inventory: nil,
	}
	return &newPlayer
}

func (p *Player) ChangeLocation(location *Node) {
	p.Location = location
}
func (p *Player) GetBackPack() {
	var newBack = Back{
		Items:    []string{},
		Capacity: 10,
	}
	var newInventory = new(Inventory)
	newInventory.Items = &newBack
	p.Inventory = newInventory
	fmt.Println("был создан инвентарь размером")
}

func (p *Player) GetPlayerInfo() {
	fmt.Printf("\n player name: %+v", p.Nickname)
	fmt.Printf("\n inventory: %v \n", p.Inventory.PrintItems())
	fmt.Printf("current location : %v \n", p.Location.Name)
}
