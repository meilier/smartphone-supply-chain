package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//HomeHandler : home page
func (app *Application) AddLogisticsHandler(w http.ResponseWriter, r *http.Request) {
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
					cn = v.UserOperation["AddLogistics"].ChannelName
					ccn = v.UserOperation["AddLogistics"].CCName
					fcn = v.UserOperation["AddLogistics"].Fcn
					break
				}
			}
			key1 := webutil.PhoneType
			key2 := r.FormValue("bnumber")
			key := key1 + key2

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
	renderTemplate(w, r, "addlogistics.html", data)
}

//HomeHandler : home page
func (app *Application) GetLogisticsHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var tInfo webutil.TransitInfo
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
					cn = v.UserOperation["GetLogistics"].ChannelName
					ccn = v.UserOperation["GetLogistics"].CCName
					fcn = v.UserOperation["GetLogistics"].Fcn
					break
				}
			}
			fmt.Println("cn,ccn,fcn GetLogistics is", cn, ccn, fcn)
			key := webutil.PhoneType + r.FormValue("bnumber")

			companyInfo, err := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
			if err != nil {
				http.Error(w, "Unable to invoke hello in the blockchain", 500)
			}
			json.Unmarshal([]byte(companyInfo), &tInfo)
			fmt.Println("logistics is cinfo is", companyInfo, tInfo)
		}
		data["LogisticsInfo"] = tInfo.ConcreteTransitInfo
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	fmt.Println("org and username", oName, uName)
	var batch []string
	if oName == "smartphone" {
		batch = app.GetPhoneBatchInfo("logistics", uName)
	} else {
		batch = app.GetBatchInfo(oName, uName)
	}
	data["BatchInfo"] = batch
	renderTemplate(w, r, "getlogistics.html", data)
}

func (app *Application) DeleteLogisticsHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	data = make(map[string]interface{})
	uName := webutil.MySession.GetUserName(r)
	oName := webutil.MySession.GetOrgName(r)
	//oName := webutil.MySession.GetOrgName(r)
	if len(uName) == 0 {
		http.Redirect(w, r, "./login.html", 302)
		return
	}
	// get supplierinfo first
	if r.FormValue("submitted") == "true" {
		//befor send request we need to check session
		uName := webutil.MySession.GetUserName(r)
		oName := webutil.MySession.GetOrgName(r)

		//add properties to args
		//TODO: here to map batchinfo to data
		bnumber := r.FormValue("bnumber")
		key := bnumber
		trinfo := app.GetPhoneLogisticsInfo(oName, uName, key)
		fmt.Println("cinfo is", trinfo)
		data["BatchNumber"] = bnumber
		data["LogisticsInfo"] = trinfo.ConcreteTransitInfo
		data["ShowInfo"] = true
		data["Hide"] = true

		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	if r.FormValue("delete") == "true" {
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
					cn = v.UserOperation["DeleteLogistics"].ChannelName
					ccn = v.UserOperation["DeleteLogistics"].CCName
					fcn = v.UserOperation["DeleteLogistics"].Fcn
					break
				}
			}
			key1 := webutil.PhoneType
			key2 := r.FormValue("bnumber")
			key := key1 + key2

			//add properties to args
			passargs = append(passargs, key)
			fmt.Println("add sub aaaa", cn, ccn, fcn)
			txid, err := fSetup.DeleteCC(cn, ccn, fcn, passargs)
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
	fmt.Println("wzx batch is", batch)
	data["BatchInfo"] = batch
	renderTemplate(w, r, "deletelogistics.html", data)
}
