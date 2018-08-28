package controllers

import (
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

func (app *Application) AddSupplierHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})

	//befor send request we need to check session
	var passargs []string
	uName := webutil.MySession.GetUserName(r)
	oName := webutil.MySession.GetOrgName(r)
	var cn string
	var ccn string
	var fcn string
	//find cfg name
	for _, v := range webutil.Orgnization[oName] {
		if v.UserName == uName {
			cn = v.UserOperation["AddSupplier"].ChannelName
			ccn = v.UserOperation["AddSupplier"].CCName
			fcn = v.UserOperation["AddSupplier"].Fcn
			break
		}
	}
	if r.FormValue("submitted") == "true" {
		uName := webutil.MySession.GetUserName(r)
		if fSetup, ok := app.Fabric[uName]; ok {
			key1 := r.FormValue("pmodel")
			key2 := r.FormValue("snumber")
			key := key1 + key2

			suppliertypeValue := r.FormValue("stype")
			nameValue := r.FormValue("sname")
			locationVaule := r.FormValue("slocation")
			//add properties to args
			passargs = append(passargs, key)
			passargs = append(passargs, suppliertypeValue)
			passargs = append(passargs, nameValue)
			passargs = append(passargs, locationVaule)

			txid, err := fSetup.InvokeCC(cn, ccn, fcn, passargs)
			if err != nil {
				http.Error(w, "Unable to invoke hello in the blockchain", 500)
			}
			data["TransactionId"] = txid
			data["Success"] = true
			data["Response"] = true
		}
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	renderTemplate(w, r, "addsupplier.html", data)
}
