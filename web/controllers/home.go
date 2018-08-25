package controllers

import (
	"net/http"
)

//HomeHandler : home page
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	supplierValue, err := app.Fabric.QuerySupplier()
	println("suppliervalue is ", supplierValue)
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["SupplierInfo"] = supplierValue

	renderTemplate(w, r, "home.html", data)
}
