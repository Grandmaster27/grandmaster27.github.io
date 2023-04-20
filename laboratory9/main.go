package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"text/template"
)

type server struct {
	db *sql.DB
}

func database() server {
	database, _ := sql.Open("sqlite3", "website_users.db")

	server := server{db: database}

	return server
}

func (server *server) new_userPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		name := r.FormValue("name")
		last_name := r.FormValue("last_name")
		age := r.FormValue("age")
		favoriteManga := r.FormValue("favoriteManga")
		_, err = server.db.Exec("INSERT INTO website_users(name, lastName, age, favoriteManga) VALUES ($1, $2, $3, $4)", name, last_name, age, favoriteManga)
		if err != nil {
			log.Fatal(err)
		}
		mapp := map[string]interface{}{"name": name, "last_name": last_name}
		tmpl, _ := template.ParseFiles("static/new_user.html")
		tmpl.Execute(w, mapp)
		return
	}
	tmpl, _ := template.ParseFiles("static/new_user.html")
	tmpl.Execute(w, nil)

}

func main() {
	server := database()

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/new_user", server.new_userPage)

	http.ListenAndServe(":8080", nil)
	defer server.db.Close()
}
