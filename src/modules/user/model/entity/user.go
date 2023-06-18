package entity

import "time"

type User struct {
	Id         int
	Name       string
	Email      string
	Password   string
	Created_At time.Time
	Updated_At time.Time
}
