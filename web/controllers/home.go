package controllers

import (
	"net/http"
)

//HomeHandler : home page
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	supplierValue, err := app.Fabric.QuerySupplier()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Supplier string
	}{
		Supplier: supplierValue,
	}
	renderTemplate(w, r, "home.html", data)
}
