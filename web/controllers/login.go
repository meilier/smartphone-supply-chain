package controllers

import (
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//LoginHandler : home page
func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	data = make(map[string]interface{})
	fmt.Println("login.html")
	if r.FormValue("submitted") == "true" {
		uname := r.FormValue("username")
		pword := r.FormValue("password")
		org := r.FormValue("org")
		println(uname, pword, org)
		//according uname, comparing pword with map[uname]
		for _, v := range webutil.Orgnization[org] {
			fmt.Println("org user", v.UserName)
			if v.UserName == uname {
				if v.Secret == pword {
					webutil.MySession.SetSession(uname, org, w)
					http.Redirect(w, r, "./home.html", 302)
					return
				}
			}

		}
		//login failed redirect to login page and show failed
		data["LoginFailed"] = true
		loginTemplate(w, r, "login.html", data)
		return
	}
	loginTemplate(w, r, "login.html", data)
}
