package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewDeleteEmployeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			id := r.URL.Query().Get("id")
			r.ParseForm()

			name := r.Form["name"][0]
			address := r.Form["address"][0]
			npwp := r.Form["npwp"][0]

			_, err := db.Exec("UPDATE employe SET name=?, npwp=?, address=? WHERE id = ?", name, npwp, address, id)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)

				return
			}

			http.Redirect(w, r, "/employe", http.StatusMovedPermanently)

		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT name, npwp, address FROM employe WHERE id = ?", id)

			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)

				return
			}

			var employe Employe
			err := row.Scan(
				&employe.Name,
				&employe.Npwp,
				&employe.Address,
			)

			employe.Id = id

			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)

				return
			}

			fp := filepath.Join("views", "update.html")

			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)

				return
			}

			data := make(map[string]any)
			data["employe"] = employe
			err = tmpl.Execute(w, data)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)

				return
			}
		}

	}
}
