package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//HomeHandler : home page
func (app *Application) AddSalesHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	data = make(map[string]interface{})
	uName := webutil.MySession.GetUserName(r)
	oName := webutil.MySession.GetOrgName(r)
	//oName := webutil.MySession.GetOrgName(r)
	if len(uName) == 0 {
		http.Redirect(w, r, "./login.html", 302)
		return
	}
	if r.FormValue("submitted") == "true" {
		uName := webutil.MySession.GetUserName(r)
		oName := webutil.MySession.GetOrgName(r)
		if fSetup, ok := app.Fabric[uName]; ok {
			//befor send request we need to check session
			var passargs []string

			var cn string
			var ccn string
			var fcn string
			//find cfg name
			for _, v := range webutil.Orgnization[oName] {
				if v.UserName == uName {
					cn = v.UserOperation["AddSales"].ChannelName
					ccn = v.UserOperation["AddSales"].CCName
					fcn = v.UserOperation["AddSales"].Fcn
					break
				}
			}
			key1 := webutil.PhoneType
			key2 := r.FormValue("snumber")
			key := key1 + key2
			// here key2 cannot bigger than max batch number
			n := r.FormValue("name")
			l := r.FormValue("location")
			m := r.FormValue("manager")
			d := r.FormValue("date")
			//add properties to args
			passargs = append(passargs, key)
			passargs = append(passargs, n)
			passargs = append(passargs, l)
			passargs = append(passargs, m)
			passargs = append(passargs, d)
			fmt.Println("cn,ccn,fcn addsales is", cn, ccn, fcn)
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
	fmt.Println("org and username", oName, uName)
	batch := app.GetBatchInfo(oName, uName)
	data["BatchInfo"] = batch
	renderTemplate(w, r, "addsales.html", data)
}

//HomeHandler : home page
func (app *Application) GetSalesHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var sInfo webutil.SalesInfo
	data = make(map[string]interface{})
	uName := webutil.MySession.GetUserName(r)
	oName := webutil.MySession.GetOrgName(r)
	//oName := webutil.MySession.GetOrgName(r)
	if len(uName) == 0 {
		http.Redirect(w, r, "./login.html", 302)
		return
	}

	if r.FormValue("submitted") == "true" {
		uName := webutil.MySession.GetUserName(r)
		oName := webutil.MySession.GetOrgName(r)
		if fSetup, ok := app.Fabric[uName]; ok {
			//befor send request we need to check session

			var cn string
			var ccn string
			var fcn string
			//find cfg name
			for _, v := range webutil.Orgnization[oName] {
				if v.UserName == uName {
					cn = v.UserOperation["GetSales"].ChannelName
					ccn = v.UserOperation["GetSales"].CCName
					fcn = v.UserOperation["GetSales"].Fcn
					break
				}
			}
			fmt.Println("cn,ccn,fcn getsales is", cn, ccn, fcn)
			key := webutil.PhoneType + r.FormValue("snumber")

			companyInfo, err := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
			if err != nil {
				http.Error(w, "Unable to invoke hello in the blockchain", 500)
			}
			json.Unmarshal([]byte(companyInfo), &sInfo)
			fmt.Println("companyInfo is cinfo is", companyInfo, sInfo)
		}
		data["SalesInfo"] = sInfo
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	fmt.Println("org and username", oName, uName)
	var batch []string
	if oName == "smartphone" {
		batch = app.GetPhoneBatchInfo("sales", uName)
	} else {
		batch = app.GetBatchInfo(oName, uName)
	}
	data["BatchInfo"] = batch
	renderTemplate(w, r, "getsales.html", data)
}
