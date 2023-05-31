package repository

import (
	"context"
	"database/sql"
	"fmt"
	"task-management-be/src/model/entity"
)

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) Create(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	insertQuery := `insert into users("name", "email", "password", "created_at", "updated_at") values ($1, $2, $3, $4, $5)`

	_, err := tx.ExecContext(ctx, insertQuery, user.Name, user.Email, user.Password, user.Created_At, user.Updated_At)

	if err != nil {
		fmt.Println("error create user:", err)
		return entity.User{}
	}

	return user
}

func (repo *UserRepository) Update(ctx context.Context, tx *sql.Tx, id string, user entity.User) entity.User {
	panic("not implemented") // TODO: Implement
}

func (repo *UserRepository) FindById(ctx context.Context, tx *sql.Tx, id string) entity.User {
	panic("not implemented") // TODO: Implement
}

func (repo *UserRepository) Delete(ctx context.Context, tx *sql.Tx, id string) bool {
	panic("not implemented") // TODO: Implement
}

func (repo *UserRepository) List(ctx context.Context, tx *sql.Tx) []entity.User {
	insertQuery := `select * from users`
	rows, err := tx.QueryContext(ctx, insertQuery)
	if err != nil {
		fmt.Println("fail")
	}

	var users []entity.User

	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Created_At, &user.Updated_At)
		if err != nil {
			fmt.Println(err)
		}

		users = append(users, user)
	}
	return users
}
