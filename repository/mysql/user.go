package mysql

import (
	"fmt"
	"quizGameGo/entities"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	panic("implement me")
}

func (d *MySQLDB) Register(u entities.User) (entities.User, error) {
	result, err := d.db.Exec(`INSERT INTO users(name, phone_number) values(?, ?)`, u.Name, u.PhoneNumber)
	if err != nil {
		return entities.User{}, fmt.Errorf("Error inserting user: %w", err)
	}

	id, _ := result.LastInsertId()
	u.ID = uint(id)

	return u, nil
}
