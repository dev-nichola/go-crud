package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteEmployeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		_, err := db.Exec("DELETE FROM employe WHERE id = ?", id)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		http.Redirect(w, r, "/employe", http.StatusMovedPermanently)
	}
}
