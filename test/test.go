package main

import (
	"fmt"
	"live/internal/entity"
	"live/internal/service"
)

func main() {
	addFakeUsers()
	addFakeRooms()
	addFakeGoods()
}

func testUserService() {
	var userService service.UserService = service.NewUserServiceImpl()

	newId := userService.AddUser(&entity.User{Name: "root"})
	fmt.Printf("Added TxID = %d\n", newId)

	user := userService.QueryUserByID(newId)
	fmt.Printf("Query1: %v\n", user)

	user.Name = "qyanzh"
	user.ID = 1000
	userService.UpdateUser(user)
	fmt.Println("Updated")
	user = userService.QueryUserByID(1000)
	fmt.Printf("Query2: %v\n", user)

	userService.DeleteUser(user)
	fmt.Println("Deleted")
	user = userService.QueryUserByID(newId)
	fmt.Printf("Query3: %v\n", user)
}

func addFakeUsers() {
	var userService service.UserService = service.NewUserServiceImpl()
	userService.AddUser(&entity.User{Name: "王雷"})
	userService.AddUser(&entity.User{Name: "朱振亿"})
	userService.AddUser(&entity.User{Name: "林北"})
	userService.AddUser(&entity.User{Name: "卢本伟"})
	userService.AddUser(&entity.User{Name: "刘华强"})
	userService.AddUser(&entity.User{Name: "路飞"})
}

func addFakeGoods() {
	var goodService service.GoodService = service.NewGoodServiceImpl()
	goodService.AddGood(&entity.Good{Name: "三文鱼"})
	goodService.AddGood(&entity.Good{Name: "四文鱼"})
	goodService.AddGood(&entity.Good{Name: "卡布奇诺"})
	goodService.AddGood(&entity.Good{Name: "电脑屏幕"})
	goodService.AddGood(&entity.Good{Name: "西瓜"})
	goodService.AddGood(&entity.Good{Name: "橡胶果实"})
}

func addFakeRooms() {
	var roomService service.RoomService = service.NewRoomServiceImpl()
	roomService.AddRoom(&entity.Room{Name: "卖鱼"})
	roomService.AddRoom(&entity.Room{Name: "卖咖啡"})
	roomService.AddRoom(&entity.Room{Name: "卖瓜"})
}
