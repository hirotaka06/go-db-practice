package main

import (
	"log"
	"testDB/models"

	"github.com/labstack/echo/v4"
)

func main() {
	// データベースに接続する
	db, err := models.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Serverの構造体を初期化する
	s := models.Server{DB: db}

	// Echoのインスタンスを作成
	e := echo.New()

	// エンドポイントを設定
	e.GET("/users", s.GetUsers)
	e.GET("/users/:id", s.GetUser)
	e.POST("/users", s.CreateUser)
	e.PUT("/users/:id", s.UpdateUserHandler)
	e.DELETE("/users/:id", s.DeleteUserHandler)

	// サーバーを開始
	e.Logger.Fatal(e.Start(":8080"))
}
