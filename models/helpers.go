package models

func CreateTestUser() *User {
	u := new(User)
	u.Username = "admin"
	u.Password = "1234"
	return u
}