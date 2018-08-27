package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//HomeHandler : home page
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	var supplierInfo webutil.CompanyInfo
	supplierValue, err := app.Fabric.QuerySupplier()
	json.Unmarshal([]byte(supplierValue), &supplierInfo)
	println("suppliervalue is ", supplierValue)
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["SupplierInfo"] = supplierInfo.ConcreteCompanyInfo

	//different nav bar for different organizations

	renderTemplate(w, r, "home.html", data)
}
