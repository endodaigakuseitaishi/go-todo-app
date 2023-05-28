package main

import (
	"fmt"
	// "log"
	"go-todo-app/app/controllers"
	"go-todo-app/app/models"
	// "go-todo-app/config"
	// "log"
)	

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("taishi.endo0616@gmail.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)

	// 変数の宣言
	var num int = 10

	// ポインタの宣言
	var ptr *int

	// ポインタに変数のアドレスを代入
	ptr = &num

	// ポインタの値を表示
	fmt.Println("ポインタの値:", ptr)

	// ポインタを経由して変数の値にアクセス
	fmt.Println("ポインタ経由の値:", *ptr)

	// 変数の値を変更
	*ptr = 20

	// 変数の値が変更されたことを確認
	fmt.Println("変更後の値:", num)
}

