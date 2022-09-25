package main

import (
	"fmt"
	"game/nurlan/Locations"
	"game/nurlan/Player"
)

/*
	код писать в этом файле
	наверняка у вас будут какие-то структуры с методами, глобальные перменные ( тут можно ), функции
*/

func main() {
	/*
		в этой функции можно ничего не писать
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
	initGame()

}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все команты
		если что-то было - оно корректно перезатирается
	*/
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)
	var player = Player.NewPlayer(name)
	player.GetPlayerInfo()
	player.GetBackPack()
	player.GetPlayerInfo()

	var graph = Locations.InitializeMap()
	var kitchen = Locations.CreateNode("kitchen")
	var hall = Locations.CreateNode("hall")
	graph.AddNode(kitchen)
	graph.AddNode(hall)
	graph.AddEdge(kitchen, hall)
	graph.PrintGraph()
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return "not implemented"
}
