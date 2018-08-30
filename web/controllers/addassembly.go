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
	if len(uName) == 0 {
		http.Redirect(w, r, "./login.html", 302)
		return
	}
	var passargs []string
	var cn string
	var ccn string
	var fcn string

	for _, v := range webutil.Orgnization[oName] {
		fmt.Println("org user", v.UserName)
		if v.UserName == uName {
			cn = v.UserOperation["AddAssembly"].ChannelName
			ccn = v.UserOperation["AddAssembly"].CCName
			fcn = v.UserOperation["AddAssembly"].Fcn
			fmt.Println("query channel is ", cn)
			break
		}
	}
	key := webutil.PhoneType + r.FormValue("bnumber")
	name := r.FormValue("name")
	location := r.FormValue("location")
	manager := r.FormValue("manager")
	date := r.FormValue("date")
	r.FormValue("bnumber")
	//add properties to args
	passargs = append(passargs, key)
	passargs = append(passargs, name)
	passargs = append(passargs, location)
	passargs = append(passargs, manager)
	passargs = append(passargs, date)

	txid, err := app.Fabric[uName].InvokeCC(cn, ccn, fcn, passargs)
	if err != nil {
		http.Error(w, "Unable to invoke hello in the blockchain", 500)
	}
	data["TransactionId"] = txid
	data["Success"] = true
	data["Response"] = true

	//different nav bar for different organizations

	renderTemplate(w, r, "home.html", data)
}

//HomeHandler : home page
func (app *Application) GetAssemblyHandler(w http.ResponseWriter, r *http.Request) {
	var assemblyInfo webutil.AssemblyInfo
	uName := webutil.MySession.GetUserName(r)
	oName := webutil.MySession.GetOrgName(r)
	if len(uName) == 0 {
		http.Redirect(w, r, "./login.html", 302)
		return
	}
	fmt.Println(len(uName))

	var cn string
	var ccn string
	var fcn string

	for _, v := range webutil.Orgnization[oName] {
		fmt.Println("org user", v.UserName)
		if v.UserName == uName {
			cn = v.UserOperation["GetAssembly"].ChannelName
			ccn = v.UserOperation["GetAssembly"].CCName
			fcn = v.UserOperation["GetAssembly"].Fcn
			fmt.Println("query channel is ", cn)
			break
		}
	}
	key := webutil.PhoneType + r.FormValue("bnumber")
	assemblyValue, err := app.Fabric[uName].QueryCC(cn, ccn, fcn, []byte(key))
	json.Unmarshal([]byte(assemblyValue), &assemblyInfo)
	println("assemblyvalue is ", assemblyValue)
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["AssemblyInfo"] = assemblyInfo

	//different nav bar for different organizations

	renderTemplate(w, r, "getassembly.html", data)
}
