package main

import (
	"log"
	"testDB/infrastructure"
	"testDB/infrastructure/repositories"
	"testDB/interfaces/handlers"
	"testDB/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	// データベースに接続する
	db, err := infrastructure.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// リポジトリ、ユースケース、ハンドラを初期化する
	userRepo := &repositories.UserRepository{DB: db}
	userUsecase := &usecases.UserUsecase{Repo: userRepo}
	userHandler := &handlers.UserHandler{Usecase: userUsecase}

	// Echoのインスタンスを作成
	e := echo.New()

	// エンドポイントを設定
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	// サーバーを開始
	e.Logger.Fatal(e.Start(":8080"))
}
