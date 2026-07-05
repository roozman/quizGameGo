package mysql

import (
	"database/sql"
	"fmt"
	"quizGameGo/entities"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entities.User{}
	var createdAt []uint8

	row := d.db.QueryRow(`SELECT * from users where phone_number = ?`, phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("cant scan query result: %w", err)
	}
	return false, nil
}

func (d *MySQLDB) Register(u entities.User) (entities.User, error) {
	result, err := d.db.Exec(`INSERT INTO users(name, phone_number, password) values(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entities.User{}, fmt.Errorf("error inserting user: %w", err)
	}

	id, _ := result.LastInsertId()
	u.ID = uint(id)

	return u, nil
}

func (d *MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entities.User, bool, error) {
	user := entities.User{}
	var createdAt []uint8

	row := d.db.QueryRow(`SELECT * from users where phone_number = ?`, phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return entities.User{}, false, nil
		}
		return entities.User{}, false, fmt.Errorf("cant scan query result: %w", err)
	}
	return user, true, nil
}
