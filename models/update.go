package models

import "fmt"

func (s *Server) UpdateUser(id int, name string) error {
	// トランザクション開始
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("transaction start error: %v", err)
	}

	// トランザクション中にエラーが発生した場合に確実にロールバックする
	defer tx.Rollback()

	// UPDATE文を実行する
	if _, err = tx.Exec("UPDATE users SET name = ? WHERE id = ?", name, id); err != nil {
		return fmt.Errorf("delete user error: %v", err)
	}

	// 成功したらコミットする
	return tx.Commit()
}
