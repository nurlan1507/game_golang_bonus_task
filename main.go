package main

import (
	"bufio"
	"fmt"
	CommandsClass "game/nurlan/Commands"
	"game/nurlan/Player"
	"os"
	"strings"
)

/*
код писать в этом файле
наверняка у вас будут какие-то структуры с методами, глобальные перменные ( тут можно ), функции
*/
var Graph = Player.InitializeMap()
var ActivePerson = Player.NewPlayer("")
var Commands = CommandsClass.Command{Player: ActivePerson, GameMap: Graph}
var Step = 0 //количество ходов который сделает игрок
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

	INITMAP()
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	ActivePerson.Nickname = name
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Введите команду: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(handleCommand(text))
		Step++
		continue
	}

}

func handleCommand(commandString string) string {
	commands := Split(commandString)
	switch commands[0] {
	case "осмотреться":
		if Step == 0 {
			return Commands.FirstLookAround()
		}
		return Commands.LookAround()
	case "идти":
		return Commands.Go(commands[1])
	case "взять":
		return Commands.TakeItem(commands[1])
	case "применить":
		return Commands.UseItem(commands[1], commands[2])
	case "надеть":
		return Commands.TakeItem(commands[1])
	}
	return "not implemented"
}

func Split(str string) []string {
	arr := strings.Split(str, " ")
	return arr
}

func INITMAP() {

	//generate map
	var kitchen = Player.CreateRoom("кухня")
	var room = Player.CreateRoom("комната")
	var hall = Player.CreateRoom("коридор")
	var outside = Player.CreateRoom("улица")
	hall.Door = Player.GenerateDoor(outside)
	Graph.AddRoom(room)
	Graph.AddRoom(kitchen)
	Graph.AddRoom(hall)
	Graph.AddRoom(outside)
	Graph.ConnectRooms(kitchen, hall)
	Graph.ConnectRooms(hall, outside)
	Graph.ConnectRooms(hall, room)
	ActivePerson.ChangeLocation(kitchen)
	Graph.PrintMap()

	//generate items
	back := Player.Back{Items: nil, Capacity: 10, Name: "рюкзак"}
	key := Player.Key{Name: "ключи"}
	conspect := Player.Conspect{Name: "конспекты"}
	tea := Player.Tea{Name: "чай"}
	room.RoomObjects["стул"] = append(kitchen.RoomObjects["стул"], &back)
	kitchen.RoomObjects["стол"] = append(kitchen.RoomObjects["стол"], &tea)
	room.RoomObjects["стол"] = append(room.RoomObjects["стол"], []Player.IItem{&key, &conspect}...)
}
