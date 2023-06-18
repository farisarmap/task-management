package repository

import (
	"context"
	"database/sql"
	"fmt"
	"task-management-be/src/modules/user/model/entity"
)

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
