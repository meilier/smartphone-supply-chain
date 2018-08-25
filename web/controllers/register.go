package controllers

import (
	"net/http"
)

//RegisterHandler : home page
func (app *Application) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})
	renderTemplate(w, r, "register.html", data)
}
