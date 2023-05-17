package main

import (
	"fmt"
	"go-todo-app/app/models"
	// "go-todo-app/config"
	// "log"
)	

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "testtestpassword2"
	// user, _ := models.GetUser(3)
	// user.CreatedTodo("todo1")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(5)
	// user.CreatedTodo("todo5")
	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(2)
	// fmt.Println(user2)
	// todos, _ := user2.GetTodosByUser()

	// for _, v := range todos {
	// 	fmt.Println(v)
// 	// }
// 	todo, _ := models.GetTodo(2)
// 	todo.Content = "update todo"
// 	todo.UpdateTodo()
t, _ := models.GetTodo(4)
fmt.Println(t)
t.DeleteTodo()
}

