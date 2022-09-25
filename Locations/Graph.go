package Locations

import "fmt"

type IGameMap interface {
	AddNode(node *Node) *Node
	AddEdge(from, to *Node)
}

type GameMap struct {
	IGameMap
	Nodes []*Node
	Edges map[Node][]*Node
}

func (g *GameMap) AddNode(node *Node) *Node {
	var contains bool = g.Contains(node)
	if contains == true {
		fmt.Printf("\n node %+v already exists \n", *node)
		return node
	}
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *GameMap) AddEdge(from, to *Node) {
	if g.Edges[*from] == nil {
		g.Edges[*from] = []*Node{to}
	} else {
		g.Edges[*from] = append(g.Edges[*from], to)
	}
}

func (g *GameMap) PrintGraph() {
	for _, v := range g.Nodes {
		fmt.Printf("%v ->", *v)
		for i := 0; i < len(g.Edges[*v]); i++ {
			fmt.Printf("%+v , ", g.Edges[*v][i].Value)
		}
		fmt.Println()
	}
}

func InitializeMap() *GameMap {
	var gameMap = GameMap{
		Nodes: []*Node{},
		Edges: map[Node][]*Node{},
	}
	return &gameMap
}

type Node struct {
	GameObjectType any
	Value          any
}

func CreateNode(value string) *Node {
	var newNode = Node{
		Value: value,
	}
	return &newNode
}

//func CreateRoom(roomName string) *Node {
//	var room = Room{
//		RoomName: roomName,
//		RoomObjects: []*RoomObject,
//	}
//	var newRoom = Node{
//		GameObjectType: room,
//	}
//	return &newRoom
//}

func (g *GameMap) Contains(node *Node) bool {
	for _, v := range g.Nodes {
		if v == node {
			return true
		}
	}
	return false
}
