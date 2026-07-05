package entities

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	// Password keeps the hashed password
	Password string
}
