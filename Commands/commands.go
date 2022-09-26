package Commands

import (
	"game/nurlan/Locations"
	"game/nurlan/Player"
)

func Go() {

}
func LookAround(gameMap *Locations.GameMap, p *Player.Player) string {
	var currentLocation = p.Location
	var availableLocations = gameMap.Edges[currentLocation]
	var str string
	for _, v := range availableLocations {
		str = str + v.Name
	}
	return str
}
