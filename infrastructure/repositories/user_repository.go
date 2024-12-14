package repositories

import (
	"database/sql"
	"fmt"
	"testDB/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetAll() ([]entities.User, error) {
	var users []entities.User
	rows, err := r.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, fmt.Errorf("scan user error: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("scan users error: %v", err)
	}

	return users, nil
}

func (r *UserRepository) GetByID(id int) (entities.User, error) {
	var user entities.User
	row := r.DB.QueryRow("SELECT * FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return entities.User{}, fmt.Errorf("scan user error: %v", err)
	}
	return user, nil
}

func (r *UserRepository) Create(user entities.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		return fmt.Errorf("insert user error: %v", err)
	}
	return nil
}

func (r *UserRepository) Update(id int, name string) error {
	_, err := r.DB.Exec("UPDATE users SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return fmt.Errorf("update user error: %v", err)
	}
	return nil
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete user error: %v", err)
	}
	return nil
}
