package domain

import "time"

type Project struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
}

type ProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

type ProjectResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
