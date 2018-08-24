package controllers

import (
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//HomeHandler : home page
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	supplierValue, err := app.Fabric.QuerySupplier()
	println("suppliervalue is ", supplierValue)
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		SupplierInfo string
		LoginStatus  bool
	}{
		SupplierInfo: supplierValue,
		LoginStatus:  false,
	}
	userName := webutil.MySession.GetUserName(r)
	if userName != "" {
		//show logout
		data.LoginStatus = true
	} else {
		http.Redirect(w, r, "/", 302)
	}
	renderTemplate(w, r, "home.html", data)
}
