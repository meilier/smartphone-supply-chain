package controllers

import (
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

func (app *Application) AddBatchHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})

	if r.FormValue("submitted") == "true" {
		//befor send request we need to check session
		uName := webutil.MySession.GetUserName(r)
		oName := webutil.MySession.GetOrgName(r)
		if fSetup, ok := app.Fabric[uName]; ok {

			var passargs []string
			var cn string
			var ccn string
			var fcn string

			suppliertypeValue := r.FormValue("suppliertype")
			fmt.Println("suppliertype is", suppliertypeValue)
			//find cfg name
			//according supplier type to choose corresponding channel
			for _, v := range webutil.Orgnization[oName] {
				if v.UserName == uName {
					switch suppliertypeValue {
					case "battery":
						cn = v.UserOperation["AddBatchBattery"].ChannelName
						ccn = v.UserOperation["AddBatchBattery"].CCName
						fcn = v.UserOperation["AddBatchBattery"].Fcn
					case "display":
						cn = v.UserOperation["AddBatchDisplay"].ChannelName
						ccn = v.UserOperation["AddBatchDisplay"].CCName
						fcn = v.UserOperation["AddBatchDisplay"].Fcn
					case "cpu":
						cn = v.UserOperation["AddBatchCpu"].ChannelName
						ccn = v.UserOperation["AddBatchCpu"].CCName
						fcn = v.UserOperation["AddBatchCpu"].Fcn
					case "assembly":
						cn = v.UserOperation["AddBatchAssembly"].ChannelName
						ccn = v.UserOperation["AddBatchAssembly"].CCName
						fcn = v.UserOperation["AddBatchAssembly"].Fcn
					case "sales":
						cn = v.UserOperation["AddBatchLogistics"].ChannelName
						ccn = v.UserOperation["AddBatchLogistics"].CCName
						fcn = v.UserOperation["AddBatchLogistics"].Fcn
					}
					break
				}
			}
			key := r.FormValue("pmodel")
			batchnumber := r.FormValue("bnumber")
			//add properties to args
			passargs = append(passargs, key)
			passargs = append(passargs, batchnumber)
			fmt.Println("cn,ccn,fcn", cn, ccn, fcn)

			txid, err := fSetup.InvokeCC(cn, ccn, fcn, passargs)
			if err != nil {
				http.Error(w, "Unable to invoke chaincode in the blockchain", 500)
			}
			data["TransactionId"] = txid
			data["Success"] = true
			data["Response"] = true
		}
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	renderTemplate(w, r, "addbatch.html", data)
}

func (app *Application) GetBatchHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	data = make(map[string]interface{})

	if r.FormValue("submitted") == "true" {
		//befor send request we need to check session
		uName := webutil.MySession.GetUserName(r)
		oName := webutil.MySession.GetOrgName(r)
		if fSetup, ok := app.Fabric[uName]; ok {

			var passargs []string
			var cn string
			var ccn string
			var fcn string

			suppliertypeValue := r.FormValue("suppliertype")
			//find cfg name
			//according supplier type to choose corresponding channel
			for _, v := range webutil.Orgnization[oName] {
				if v.UserName == uName {
					switch suppliertypeValue {
					case "Battery":
						cn = v.UserOperation["GetBatchBattery"].ChannelName
						ccn = v.UserOperation["GetBatchBattery"].CCName
						fcn = v.UserOperation["GetBatchBattery"].Fcn
					case "Display":
						cn = v.UserOperation["GetBatchDisplay"].ChannelName
						ccn = v.UserOperation["GetBatchDisplay"].CCName
						fcn = v.UserOperation["GetBatchDisplay"].Fcn
					case "Cpu":
						cn = v.UserOperation["GetBatchCpu"].ChannelName
						ccn = v.UserOperation["GetBatchCpu"].CCName
						fcn = v.UserOperation["GetBatchCpu"].Fcn
					case "Assembly":
						cn = v.UserOperation["GetBatchAssembly"].ChannelName
						ccn = v.UserOperation["GetBatchAssembly"].CCName
						fcn = v.UserOperation["GetBatchAssembly"].Fcn
					case "Sales":
						cn = v.UserOperation["GetBatchLogistics"].ChannelName
						ccn = v.UserOperation["GetBatchLogistics"].CCName
						fcn = v.UserOperation["GetBatchLogistics"].Fcn
					}
					break
				}
			}
			key := r.FormValue("pmodel")
			//add properties to args
			passargs = append(passargs, key)
			//TODO: here to map batchinfo to data
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
	renderTemplate(w, r, "getbatch.html", data)
}

func (app *Application) AddSupplierHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})

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
					cn = v.UserOperation["AddSupplier"].ChannelName
					ccn = v.UserOperation["AddSupplier"].CCName
					fcn = v.UserOperation["AddSupplier"].Fcn
					break
				}
			}
			key1 := webutil.PhoneType
			key2 := r.FormValue("bnumber")
			key := key1 + key2

			sn := r.FormValue("supplierName")
			sl := r.FormValue("supplierLocation")
			ci := r.FormValue("cInformation")
			//add properties to args
			passargs = append(passargs, key)
			passargs = append(passargs, sn)
			passargs = append(passargs, sl)
			passargs = append(passargs, ci)

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

func (app *Application) GetSupplierHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})

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
					cn = v.UserOperation["GetSupplier"].ChannelName
					ccn = v.UserOperation["GetSupplier"].CCName
					fcn = v.UserOperation["GetSupplier"].Fcn
					break
				}
			}
			key := webutil.PhoneType + r.FormValue("bnumber")

			//add properties to args
			passargs = append(passargs, key)

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

func (app *Application) AddSubComponetHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})

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
					cn = v.UserOperation["AddCompanyInfo"].ChannelName
					ccn = v.UserOperation["AddCompanyInfo"].CCName
					fcn = v.UserOperation["AddCompanyInfo"].Fcn
					break
				}
			}
			key := webutil.PhoneType + r.FormValue("bnumber")

			sn := r.FormValue("subcomponetName")
			sl := r.FormValue("companyName")
			ci := r.FormValue("location")
			//add properties to args
			passargs = append(passargs, key)
			passargs = append(passargs, sn)
			passargs = append(passargs, sl)
			passargs = append(passargs, ci)

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

func (app *Application) GetSubcomponentHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})

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
					cn = v.UserOperation["GetCompanyInfo"].ChannelName
					ccn = v.UserOperation["GetCompanyInfo"].CCName
					fcn = v.UserOperation["GetCompanyInfo"].Fcn
					break
				}
			}
			key := webutil.PhoneType + r.FormValue("bnumber")

			sn := r.FormValue("supplierName")
			sl := r.FormValue("supplierLocation")
			ci := r.FormValue("cInformation")
			//add properties to args
			passargs = append(passargs, key)
			passargs = append(passargs, sn)
			passargs = append(passargs, sl)
			passargs = append(passargs, ci)

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

func (app *Application) AddAssemblyhHandler(w http.ResponseWriter, r *http.Request) {

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
