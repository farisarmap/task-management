package middlewares

import (
	"context"
	"database/sql"
	"net/http"
	"task-management-be/src/helper"
)

func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

func DbTransactionMiddleware(handler http.Handler, db *sql.DB, context context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txHandle, err := db.BeginTx(context, nil)
		helper.ErrorNotNil(err, "Failed to start transaction db")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

	})
}
