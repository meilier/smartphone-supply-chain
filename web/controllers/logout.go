package controllers

import (
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//LogoutHandler : logout user
func (app *Application) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	webutil.MySession.ClearSession(w)
	http.Redirect(w, r, "/home.html", 302)

}
