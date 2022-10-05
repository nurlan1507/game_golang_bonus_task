package Commands

import (
	"errors"
	"fmt"
	"game/nurlan/Player"
)

type Command struct {
	GameMap *Player.GameMap
	Player  *Player.Player
}

func (c *Command) LookAround() string {
	var currentLocation = c.Player.Location
	var availableLocations = c.GameMap.Edges[currentLocation]
	//var RoomObjects map[string]*Player.Item = currentLocation.RoomObjects
	var availablePaths string
	for _, v := range availableLocations {
		availablePaths = availablePaths + v.Name + ", "
	}
	var availableRoomObjectsString string = ""
	for k, v := range currentLocation.RoomObjects {
		var objects string = ""
		for _, v := range v {
			objects += v.GetItemName() + ", "
		}
		availableRoomObjectsString = availableRoomObjectsString + fmt.Sprintf("на %ve: %v", k, objects)
	}

	return availableRoomObjectsString + "можно пройти - " + availablePaths
}

func (c *Command) FirstLookAround() string {
	var currentLocation = c.Player.Location
	var availableLocations = c.GameMap.Edges[currentLocation]

	//var RoomObjects map[string]*Player.Item = currentLocation.RoomObjects
	var availablePaths string
	for _, v := range availableLocations {
		availablePaths = availablePaths + v.Name + ", "
	}
	var availableRoomObjectsString string = ""
	currentLocationString := fmt.Sprintf("ты находишься на %vе, ", c.Spragat())
	for k, v := range currentLocation.RoomObjects {
		var objects string = ""
		for _, v := range v {
			objects += ", " + v.GetItemName()
		}
		availableRoomObjectsString = availableRoomObjectsString + fmt.Sprintf("на %ve: %v", k, objects) + " "
	}
	if availableRoomObjectsString == "" {
		return "пустая комната. можно пройти - " + availablePaths
	}
	return currentLocationString + availableRoomObjectsString + ", надо собрать рюкзак и идти в универ. " + "можно пройти - " + availablePaths
}

func (c *Command) GoHome() string {
	var availableLocations = c.GameMap.Edges[c.Player.Location]
	if len(availableLocations) == 0 {
		return "нет пути домой"
	}
	var home = availableLocations[0]
	c.Player.Location = home
	return c.DefineWhatInRoom()
}

func (c *Command) Go(goTo string) string {
	if goTo == "домой" {
		if c.Player.Location.Name != "улица" {
			return "некуда идти"
		}
		return c.GoHome()
	}
	var availableLocations = c.GameMap.Edges[c.Player.Location]
	var destinationNode *Player.Node = nil
	for _, v := range availableLocations {
		if v.Name == goTo {
			destinationNode = v
		}
	}
	if destinationNode == nil {
		return fmt.Sprintf("нет пути в %v", goTo)
	}
	if destinationNode.Name == "улица" {
		if c.CheckDoor() == true {
			return fmt.Sprintf("на улице весна, можно пройти - домой")
		}
		return "дверь закрыта"
	}
	c.Player.Location = destinationNode

	return c.DefineWhatInRoom()
}

func (c *Command) DefineWhatInRoom() string {

	var availableLocations = c.GameMap.Edges[c.Player.Location]
	var availableLocationsString string
	for _, v := range availableLocations {
		availableLocationsString += v.Name + " ,"
	}

	if len(c.Player.Location.RoomObjects) == 0 {
		return "ничего интересного. можно пройти - " + availableLocationsString + "."
	}
	return fmt.Sprintf("ты в своей %v, можно пройти - %v", c.Player.Location.Name, availableLocationsString)
}

func (c *Command) TakeItem(itemName string) string {
	if c.Player.Inventory == nil && itemName != "рюкзак" {
		return "некуда класть"
	}
	var desiredItem, err = c.FindItem(c.Player.Location, itemName)
	if err != nil {
		return err.Error()
	}

	return desiredItem.ItemInitAction(c.Player)
}

func (c *Command) FindItem(location *Player.Node, itemName string) (Player.IItem, error) {
	var desiredItem Player.IItem = nil
	if c.Player.Location != location {
		return nil, errors.New("нет такого")
	}
	var RoomObjectKey string
	for k, _ := range c.Player.Location.RoomObjects {
		RoomObjectKey = k
		for _, v := range c.Player.Location.RoomObjects[RoomObjectKey] {
			if itemName == v.GetItemName() {
				desiredItem = v
				c.RemoveItem(RoomObjectKey, itemName)
				return desiredItem, nil
			}
		}
	}
	return nil, errors.New("нет такого")
}

func (c *Command) UseItem(itemName string, useTo string) string {
	var desiredItem Player.IItem = nil
	if c.Player.Inventory == nil {
		return "нет инвентаря"
	}
	fmt.Printf("%+v", c.Player.Inventory.Items)

	for _, v := range c.Player.Inventory.Items {
		if itemName == v.GetItemName() {
			fmt.Println(v.GetItemName())
			desiredItem = v
		}
	}
	if desiredItem == nil {
		var e string = fmt.Sprintf("нет предмета в инвентаре - %v", itemName)
		return e
	}
	return desiredItem.ItemAction(c.Player)

}

func (c *Command) getBack() {

}

func (c *Command) RemoveItem(key string, itemName string) {
	var index int
	for i, v := range c.Player.Location.RoomObjects[key] {
		if v.GetItemName() == itemName {
			index = i
		}
	}
	c.Player.Location.RoomObjects[key][index] = c.Player.Location.RoomObjects[key][len(c.Player.Location.RoomObjects[key])-1]
	c.Player.Location.RoomObjects[key] = c.Player.Location.RoomObjects[key][:len(c.Player.Location.RoomObjects[key])-1]
}

func (c *Command) CheckDoor() bool {
	fmt.Println(c.Player.Location.Name)
	if c.Player.Location.Door != nil {
		if c.Player.Location.Door.Status == true {
			return true
		}
		return false
	}
	return false
}

func (c *Command) Spragat() string {
	var str = c.Player.Location.Name
	newStr := str[:(len(str) - 2)]
	return newStr
}
