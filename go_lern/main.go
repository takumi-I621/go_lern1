package main

import "go_lern/app/controllers"

func main() {
	//fmt.Println(models.Db)

	controllers.StartMainServer()

	//トレーニングレコード作成サンプル
	/*
		t := &models.Training{}
		t.Tday = "2021/03/17"
		t.Hukkin = 50
		t.Udetate = 50
		t.Haikin = 50
		t.Working = "1.5km"

		fmt.Println(t)
		t.CreteTraining()
	*/
	/*
		user, _ := models.GetUserByEmail("test@example.com")
		fmt.Println(user)

		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(session)

		valid, _ := session.CheckSession()
		fmt.Println(valid)
	*/

	/*
		//ユーザレコード作成
		u := &models.User{}
		u.Name = "test6"
		u.Email = "test6@example.com"
		u.PassWord = "test"
		fmt.Println(u)

		u.CreateUser()
	*/
	/*
		//ユーザ取得
		u, _ := models.GetUser(1)
		//取得したユーザの表示
		fmt.Println(u)
		//ユーザ情報を更新

		u.Name = "Test2"
		u.Email = "test123@example.com"

		u.UpdateUser()
		//ユーザ取得
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/
	/*
		//todoUser作成
		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo")
	*/
	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/
	/*
		user, _ := models.GetUser(3)
		user.CreateTodo("Third Todo")
	*/
	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		t, _ := models.GetTodo(3)
		t.DeleteTodo()
	*/
}
