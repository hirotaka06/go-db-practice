package models

import "fmt"

func (s *Server) InsertUsers(users []User) error {
	// トランザクション開始
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("transaction start error: %v", err)
	}

	// トランザクション中にエラーが発生した場合に確実にロールバックする
	defer tx.Rollback()

	// プリペアードステートメントを作成
	stmt, err := tx.Prepare("INSERT INTO users (name, age) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("prepared statement error: %v", err)
	}
	// 最後にプリペアードステートメントを閉じる
	defer stmt.Close()

	// userを一人一人Insertする
	for _, user := range users {
		// INSERT文を実行する
		if _, err := stmt.Exec(user.Name, user.Age); err != nil {
			return fmt.Errorf("insert user error: %v", err)
		}
	}

	// 成功したらコミットする
	return tx.Commit()
}
