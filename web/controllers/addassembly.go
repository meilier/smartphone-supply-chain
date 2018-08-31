package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//HomeHandler : home page
func (app *Application) AddAssemblyHandler(w http.ResponseWriter, r *http.Request) {
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
					cn = v.UserOperation["AddAssembly"].ChannelName
					ccn = v.UserOperation["AddAssembly"].CCName
					fcn = v.UserOperation["AddAssembly"].Fcn
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
	renderTemplate(w, r, "addassembly.html", data)
}

//HomeHandler : home page
func (app *Application) GetAssemblyHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var aInfo webutil.AssemblyInfo
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
					cn = v.UserOperation["GetAssembly"].ChannelName
					ccn = v.UserOperation["GetAssembly"].CCName
					fcn = v.UserOperation["GetAssembly"].Fcn
					break
				}
			}
			fmt.Println("cn,ccn,fcn GetAssembly is", cn, ccn, fcn)
			key := webutil.PhoneType + r.FormValue("bnumber")

			companyInfo, err := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
			if err != nil {
				http.Error(w, "Unable to invoke hello in the blockchain", 500)
			}
			json.Unmarshal([]byte(companyInfo), &aInfo)
			fmt.Println("companyInfo is cinfo is", companyInfo, aInfo)
		}
		data["AssemblyInfo"] = aInfo
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	fmt.Println("org and username", oName, uName)
	var batch []string
	if oName == "smartphone" {
		batch = app.GetPhoneBatchInfo("assembly", uName)
	} else {
		batch = app.GetBatchInfo(oName, uName)
	}

	data["BatchInfo"] = batch
	renderTemplate(w, r, "getassembly.html", data)
}
