package models

import (
	"fmt"
)

// Userの構造体
type User struct {
	ID   int
	Name string
	Age  int
}

func (s *Server) SelectUsers() ([]User, error) {
	// 構造体Userのスライス型、users
	var users []User

	// SELECT文を実行する
	rows, err := s.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}

	for rows.Next() {
		var user User
		// データベースから読み取られた列を、
		//sqlパッケージで提供されている、一般的なGoの型に変換する
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, fmt.Errorf("scan the user error: %v", err)
		}
		// Usersに追加する
		users = append(users, user)
	}

	// for文でエラーが発生した場合に呼び出される
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("scan users error: %v", err)
	}

	return users, nil
}


func (s *Server) SelectUsersByName(name string) ([]User, error) {
	// 構造体Userのスライス型、users
	var users []User

	// SELECT文を実行する
	rows, err := s.DB.Query("SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}

	for rows.Next() {
		var user User
		// データベースから読み取られた列を、
		//sqlパッケージで提供されている、一般的なGoの型に変換する
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, fmt.Errorf("scan user error: %v", err)
		}
		// Usersに追加する
		users = append(users, user)
	}

	// for文でエラーが発生した場合に呼び出される
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("scan users error: %v", err)
	}

	return users, nil
}

func (s *Server) SelectUserByID(id int) (User, error) {
	var user User

	// SELECT文を実行する
	row := s.DB.QueryRow("SELECT * FROM users WHERE id = ?", id)

	// データベースから読み取られた列を、
	//sqlパッケージで提供されている、一般的なGoの型に変換する
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, fmt.Errorf("scan user error: %v", err)
	}

	return user, nil
}
