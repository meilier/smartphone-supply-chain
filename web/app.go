package web

import (
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/controllers"
)

func Serve(app *controllers.Application) {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/home.html", app.HomeHandler)
	http.HandleFunc("/addsupplier.html", app.AddSupplierHandler)
	http.HandleFunc("/register.html", app.RegisterHandler)
	http.HandleFunc("/login.html", app.LoginHandler)
	http.HandleFunc("/logout.html", app.LogoutHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)

	// open other services at the same time, such as Store listening at 3003
}
