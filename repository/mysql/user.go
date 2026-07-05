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
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("cant scan query result: %w", err)
	}
	return false, nil
}

func (d *MySQLDB) Register(u entities.User) (entities.User, error) {
	result, err := d.db.Exec(`INSERT INTO users(name, phone_number) values(?, ?)`, u.Name, u.PhoneNumber)
	if err != nil {
		return entities.User{}, fmt.Errorf("error inserting user: %w", err)
	}

	id, _ := result.LastInsertId()
	u.ID = uint(id)

	return u, nil
}
