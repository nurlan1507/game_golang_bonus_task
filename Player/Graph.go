package Player

import (
	"fmt"
)

type IGameMap interface {
	AddRoom(node *Node) *Node
	AddEdge(from, to *Node)
}

type GameMap struct {
	IGameMap
	Nodes []*Node
	Edges map[*Node][]*Node
}
type Node struct {
	Name           string
	GameObjectType any
	RoomObjects    map[string][]IItem
	Door           *Door
}

func (g *GameMap) AddRoom(node *Node) *Node {

	var contains bool = g.Contains(node)

	if contains == true {
		fmt.Printf("\n node %+v already exists \n", *node)
		return node
	}
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *GameMap) AddEdge(from *Node, to *Node) {
	if g.Edges[from] == nil {
		g.Edges[from] = []*Node{to}
	} else {
		g.Edges[from] = append(g.Edges[from], to)
	}
}

func (g *GameMap) ConnectRooms(from *Node, to *Node) {
	g.AddEdge(from, to)
	g.AddEdge(to, from)
}

func (g *GameMap) PrintMap() {
	for _, v := range g.Nodes {
		fmt.Printf("%v ->", v.Name)
		for i := 0; i < len(g.Edges[v]); i++ {
			fmt.Printf("%+v , ", g.Edges[v][i].Name)
		}
		fmt.Println()
	}
}

func InitializeMap() *GameMap {
	var gameMap = GameMap{
		Nodes: []*Node{},
		Edges: map[*Node][]*Node{},
	}
	return &gameMap
}

//func CreateNode(node *Node) *Node {
//	var newNode = Node{
//		GameObject:node,
//	}
//	return &newNode
//}

func (g *GameMap) Contains(node *Node) bool {
	for _, v := range g.Nodes {
		if v == node {
			return true
		}
	}
	return false
}
