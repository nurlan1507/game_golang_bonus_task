package Player

import (
	"fmt"
)

type IPlayer interface {
	ChangeLocation()
}

//
//func (I *Inventory) PrintItems() string {
//	if I != nil {
//		return I.Items.PrintAll()
//	}
//	return "no backpack"
//}

type Player struct {
	Nickname  string
	Inventory *Back
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
func (p *Player) GetBackPack(newBack *Back) {
	p.Inventory = newBack
	fmt.Println(p.Inventory.Items)
	fmt.Printf("был создан инвентарь размером %v", p.Inventory.Capacity)
}

func (p *Player) GetPlayerInfo() {
	fmt.Printf("\n player name: %+v", p.Nickname)
	//fmt.Printf("\n inventory: %v \n", p.Inventory.PrintItems())
	fmt.Printf("current location : %v \n", p.Location.Name)
}
