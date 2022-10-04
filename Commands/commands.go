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
			objects += v.ItemName + ", "
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
		availablePaths = availablePaths + ", " + v.Name
	}
	var availableRoomObjectsString string = ""
	currentLocationString := fmt.Sprintf("ты находишься на %vе ,", currentLocation.Name)
	for k, v := range currentLocation.RoomObjects {
		var objects string = ""
		for _, v := range v {
			objects += v.ItemName + ", "
		}
		availableRoomObjectsString = availableRoomObjectsString + fmt.Sprintf("на %ve: %v", k, objects)
	}
	return currentLocationString + availableRoomObjectsString + ",надо собрать рюкзак и идти в универ. " + "можно пройти - " + availablePaths
}

func (c *Command) Go(goTo string) string {
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
	var desiredItem, err = c.FindItem(c.Player.Location, itemName)
	if err != nil {
		return err.Error()
	}
	return desiredItem.ItemInstance.ItemInitAction(c.Player)
}

func (c *Command) FindItem(location *Player.Node, itemName string) (*Player.Item, error) {
	var desiredItem *Player.Item = nil
	if c.Player.Location != location {
		return nil, errors.New("нет такого")
	}
	var RoomObjectKey string
	for k, _ := range c.Player.Location.RoomObjects {
		RoomObjectKey = k
	}
	for _, v := range c.Player.Location.RoomObjects[RoomObjectKey] {
		if itemName == v.ItemName {
			desiredItem = v
		}
	}
	if desiredItem == nil {
		return nil, errors.New("нет такого")
	}
	return desiredItem, nil
}

func (c *Command) getBack() {

}
