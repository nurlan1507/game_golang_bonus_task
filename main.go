package main

import (
	"fmt"
	CommandsClass "game/nurlan/Commands"
	"game/nurlan/Player"
	"strings"
)

/*
код писать в этом файле
наверняка у вас будут какие-то структуры с методами, глобальные перменные ( тут можно ), функции
*/
var Graph = Player.InitializeMap()
var ActivePerson = Player.NewPlayer("")
var Commands = CommandsClass.Command{Player: ActivePerson, GameMap: Graph}

//var ActivePlayer = Player.NewPlayer()

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

	var kitchen = Player.CreateRoom("кухня", nil)
	var hall = Player.CreateRoom("коридор", nil)
	Graph.AddRoom(kitchen)
	Graph.AddRoom(hall)
	Graph.AddEdge(kitchen, hall)
	Graph.PrintMap()
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(name)
	ActivePerson.Nickname = name
	//ActivePerson.ChangeLocation(kitchen)
	//ActivePerson.GetBackPack()
	//ActivePerson.GetPlayerInfo()
	//
	for true {
		var commandString string
		_, err := fmt.Scan(&commandString)
		command := strings.Split(commandString, " ")
		fmt.Println(command)
		if err != nil {
			fmt.Println(err)
		}
		handleCommand(commandString)
		continue
	}

}

func handleCommand(commandString string) string {
	command := strings.Split(commandString, " ")
	switch command[0] {
	case "осмотреться":
		fmt.Println(Commands.LookAround())
	case "идти":
		fmt.Println(Commands.Go(command[1]))
	}
	return "not implemented"
}
