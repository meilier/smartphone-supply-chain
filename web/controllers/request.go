package controllers

import (
	"net/http"
)

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	var passargs []string
	if r.FormValue("submitted") == "true" {
		nameValue := r.FormValue("cname")
		locationVaule := r.FormValue("clocation")
		passargs = append(passargs, nameValue)
		passargs = append(passargs, locationVaule)
		txid, err := app.Fabric.InvokeSupplier(passargs)
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "request.html", data)
}
