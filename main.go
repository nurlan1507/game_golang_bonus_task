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

	//создаю мир
	var graph = Locations.InitializeMap()
	var kitchen = Locations.CreateRoom("Kitchen")
	var hall = Locations.CreateRoom("hall")
	graph.AddRoom(kitchen)
	fmt.Printf("%T\n", graph)
	graph.AddRoom(hall)
	graph.AddRoom(hall)
	graph.AddEdge(kitchen, hall)
	graph.PrintMap()
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)
	var player = Player.NewPlayer(name, kitchen)
	player.GetPlayerInfo()
	player.GetBackPack()
	player.GetPlayerInfo()
	//

	for true {
		var command string
		var property string
		fmt.Println("Введите команду")
		_, err := fmt.Scan(&command, &property)
		if err != nil {
			fmt.Println(err)
		}
		handleCommand(command)
		continue
	}

}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	switch command {
	case "идти":

	}
	fmt.Println()
	fmt.Printf("Твоя комманда %v", command)
	fmt.Println()
	return "not implemented"
}
