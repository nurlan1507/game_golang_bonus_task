package Commands

import (
	"game/nurlan/Player"
)

type Command struct {
	GameMap *Player.GameMap
	Player  *Player.Player
}

func (c *Command) LookAround() string {
	var currentLocation = c.Player.Location
	var availableLocations = c.GameMap.Edges[currentLocation]
	var availablePaths string
	for _, v := range availableLocations {
		availablePaths = availablePaths + v.Name + ", "
	}
	switch currentLocation.Name {
	case "кухня":
		return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - " + availablePaths
	}
	return "ты находишься на " + currentLocation.Name + ""
}

func (c *Command) Go(goTo string) string {
	//проверить если есть эта комната
	var availableLocation = c.GameMap.Edges[c.Player.Location]
	var destinationNode *Player.Node = nil
	for _, v := range availableLocation {
		if v.Name == goTo {
			destinationNode = v
		}
	}
	if destinationNode == nil {
		return "нет пути в " + destinationNode.Name
	}
	c.Player.Location = destinationNode
	return ""
}
