package repository

import (
	"database/sql"
	"fmt"
	"go-crud-api/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT id, first_name, last_name, balance FROM users"
	rows, err := ur.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}
	var userList []model.User
	var userObj model.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.FirstName,
			&userObj.LastName,
			&userObj.Balance,
		)

		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	query, err := ur.connection.Prepare("INSERT INTO users(first_name, last_name, balance) " +
		"VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.FirstName, user.LastName, user.Balance).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (ur *UserRepository) GetUserByID(user_id int) (*model.User, error) {
	query, err := ur.connection.Prepare("SELECT * FROM user WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user model.User
	err = query.QueryRow(user_id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Balance,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &user, nil
}
