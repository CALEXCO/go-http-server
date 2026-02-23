package models

type DbUsers struct {
	Users []User
}

type User struct {
	Name    string `json:"name" binding:"required"`
	Mail    string `json:"mail" binding:"required"`
	Company string `json:"company" binding:"required"`
	DNI     string `json:"dni" binfing:"required"`
}

func NewUser(name string, mail string, company string, dni string) User {
	return User{
		Name:    name,
		Mail:    mail,
		Company: company,
		DNI:     dni,
	}
}

func (u *User) isEqual(user User) bool {
	return u.DNI == user.DNI
}

func (db DbUsers) userExists(u User) bool {
	for _, user := range db.Users {
		if user.isEqual(u) {
			return true
		}
	}
	return false
}

func (db DbUsers) AddNewUser(u User) bool {
	if db.userExists(u) {
		return false
	}

	db.Users = append(db.Users, u)
	return true
}
