package controllers

import (
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//LoginHandler : home page
func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
	}{}
	if r.FormValue("submitted") == "true" {
		uname := r.FormValue("username")
		pword := r.FormValue("password")
		org := r.FormValue("org")
		println(uname, pword, org)
		if uname == "wzx" && pword == "arclabw401wzx" {
			//login successfully set session and redirect to home page
			//登录成功设置session
			webutil.MySession.SetSession(uname, w)
			renderTemplate(w, r, "home.html", data)
		} else {
			//login failed redirect to login page and show failed
			renderTemplate(w, r, "login.html", data)
		}

	}
	renderTemplate(w, r, "login.html", data)
}
