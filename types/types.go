package types


import (
	"time"

)

type UserStore interface {
	GetUserByMail(mail string) (*User, error)
}

type User struct {
	ID        int    	`json:"id"`
	FirstName string 	`json:"first_name"`
	LastName  string 	`json:"last_name"`
	Email     string 	`json:"email"`
	Password  string 	`json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserPayLoad struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}