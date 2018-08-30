package controllers

import (
	"net/http"
)

//QueryPhoneHandler : home page
func (app *Application) QueryPhoneHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})
	queryphoneTemplate(w, r, "queryphoneinfo.html", data)
}
