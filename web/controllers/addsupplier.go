package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

func (app *Application) AddBatchHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	data = make(map[string]interface{})
	var bInfo webutil.BatchInfo

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
			stype := r.FormValue("suppliertype")
			//add batch must greater than max batch number
			fmt.Println("oooooname uuuuuuname is", stype, uName)
			var cnb string
			var ccnb string
			var fcnb string

			//find cfg name
			//according supplier type to choose corresponding channel
			for _, v := range webutil.Orgnization[oName] {
				if v.UserName == uName {
					switch suppliertypeValue {
					case "battery":
						cnb = v.UserOperation["GetBatchBattery"].ChannelName
						ccnb = v.UserOperation["GetBatchBattery"].CCName
						fcnb = v.UserOperation["GetBatchBattery"].Fcn
					case "display":
						cnb = v.UserOperation["GetBatchDisplay"].ChannelName
						ccnb = v.UserOperation["GetBatchDisplay"].CCName
						fcnb = v.UserOperation["GetBatchDisplay"].Fcn
					case "cpu":
						cnb = v.UserOperation["GetBatchCpu"].ChannelName
						ccnb = v.UserOperation["GetBatchCpu"].CCName
						fcnb = v.UserOperation["GetBatchCpu"].Fcn
					case "assembly":
						cnb = v.UserOperation["GetBatchAssembly"].ChannelName
						ccnb = v.UserOperation["GetBatchAssembly"].CCName
						fcnb = v.UserOperation["GetBatchAssembly"].Fcn
					case "sales":
						cnb = v.UserOperation["GetBatchLogistics"].ChannelName
						ccnb = v.UserOperation["GetBatchLogistics"].CCName
						fcnb = v.UserOperation["GetBatchLogistics"].Fcn
					}
					break
				}
			}
			//add properties to args
			//TODO: here to map batchinfo to data
			batchinfo, err := fSetup.QueryCC(cnb, ccnb, fcnb, []byte(key))
			if err != nil {
				http.Error(w, "Unable to invoke hello in the blockchain", 500)
			}
			json.Unmarshal([]byte(batchinfo), &bInfo)

			if len(bInfo.Batch) != 0 && batchnumber <= bInfo.Batch[len(bInfo.Batch)-1] {
				data["CompareFailed"] = true
				renderTemplate(w, r, "addbatch.html", data)
				return
			}

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

//to do get smartphone batch
func (app *Application) GetBatchHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	data = make(map[string]interface{})
	var bInfo webutil.BatchInfo
	if r.FormValue("submitted") == "true" {
		//befor send request we need to check session
		uName := webutil.MySession.GetUserName(r)
		oName := webutil.MySession.GetOrgName(r)
		if fSetup, ok := app.Fabric[uName]; ok {

			var cn string
			var ccn string
			var fcn string

			suppliertypeValue := r.FormValue("suppliertype")
			//find cfg name
			//according supplier type to choose corresponding channel
			for _, v := range webutil.Orgnization[oName] {
				if v.UserName == uName {
					switch suppliertypeValue {
					case "battery":
						cn = v.UserOperation["GetBatchBattery"].ChannelName
						ccn = v.UserOperation["GetBatchBattery"].CCName
						fcn = v.UserOperation["GetBatchBattery"].Fcn
					case "display":
						cn = v.UserOperation["GetBatchDisplay"].ChannelName
						ccn = v.UserOperation["GetBatchDisplay"].CCName
						fcn = v.UserOperation["GetBatchDisplay"].Fcn
					case "cpu":
						cn = v.UserOperation["GetBatchCpu"].ChannelName
						ccn = v.UserOperation["GetBatchCpu"].CCName
						fcn = v.UserOperation["GetBatchCpu"].Fcn
					case "assembly":
						cn = v.UserOperation["GetBatchAssembly"].ChannelName
						ccn = v.UserOperation["GetBatchAssembly"].CCName
						fcn = v.UserOperation["GetBatchAssembly"].Fcn
					case "sales":
						cn = v.UserOperation["GetBatchLogistics"].ChannelName
						ccn = v.UserOperation["GetBatchLogistics"].CCName
						fcn = v.UserOperation["GetBatchLogistics"].Fcn
					}
					break
				}
			}
			key := r.FormValue("pmodel")
			//add properties to args
			//TODO: here to map batchinfo to data
			batchinfo, err := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
			if err != nil {
				http.Error(w, "Unable to invoke hello in the blockchain", 500)
			}
			json.Unmarshal([]byte(batchinfo), &bInfo)
			data["PhoneModel"] = key
			data["BatchInfo"] = bInfo.Batch
		}
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	renderTemplate(w, r, "getbatch.html", data)
}

func (app *Application) GetBatchInfo(suppliertype string, username string) []string {
	var bInfo webutil.BatchInfo
	//befor send request we need to check session
	if fSetup, ok := app.Fabric[username]; ok {
		var cn string
		var ccn string
		var fcn string
		fmt.Println("username is", username)

		//find cfg name
		//according supplier type to choose corresponding channel
		for _, v := range webutil.Orgnization[suppliertype] {
			fmt.Println("v.Username is", v.UserName)
			if v.UserName == username {
				fmt.Println("aaaaaa suppliertype is ", suppliertype)
				switch suppliertype {
				case "battery":
					cn = v.UserOperation["GetBatchBattery"].ChannelName
					ccn = v.UserOperation["GetBatchBattery"].CCName
					fcn = v.UserOperation["GetBatchBattery"].Fcn
				case "display":
					cn = v.UserOperation["GetBatchDisplay"].ChannelName
					ccn = v.UserOperation["GetBatchDisplay"].CCName
					fcn = v.UserOperation["GetBatchDisplay"].Fcn
					fmt.Println("whyaaaa", v.UserOperation["GetBatchDisplay"])
				case "cpu":
					cn = v.UserOperation["GetBatchCpu"].ChannelName
					ccn = v.UserOperation["GetBatchCpu"].CCName
					fcn = v.UserOperation["GetBatchCpu"].Fcn
				case "assembly":
					cn = v.UserOperation["GetBatchAssembly"].ChannelName
					ccn = v.UserOperation["GetBatchAssembly"].CCName
					fcn = v.UserOperation["GetBatchAssembly"].Fcn
				case "logistics":
					cn = v.UserOperation["GetBatchLogistics"].ChannelName
					ccn = v.UserOperation["GetBatchLogistics"].CCName
					fcn = v.UserOperation["GetBatchLogistics"].Fcn
				case "sales":
					cn = v.UserOperation["GetBatchSales"].ChannelName
					ccn = v.UserOperation["GetBatchSales"].CCName
					fcn = v.UserOperation["GetBatchSales"].Fcn
				}
				break
			}
		}
		//add properties to args
		//TODO: here to map batchinfo to data
		fmt.Println("cn ccn fcn is", cn, ccn, fcn)
		batchinfo, _ := fSetup.QueryCC(cn, ccn, fcn, []byte("Aphone"))
		json.Unmarshal([]byte(batchinfo), &bInfo)
	}
	if len(bInfo.Batch) == 0 {
		return []string{}
	}
	return bInfo.Batch[1:]
}

func (app *Application) AddCompanyHandler(w http.ResponseWriter, r *http.Request) {

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
					cn = v.UserOperation["AddCompanyInfo"].ChannelName
					ccn = v.UserOperation["AddCompanyInfo"].CCName
					fcn = v.UserOperation["AddCompanyInfo"].Fcn
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
	fmt.Println("org and username", oName, uName)
	batch := app.GetBatchInfo(oName, uName)
	data["BatchInfo"] = batch
	renderTemplate(w, r, "addcompany.html", data)
}

func (app *Application) GetCompanyHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	var cInfo webutil.CompanyInfo
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
					cn = v.UserOperation["GetCompanyInfo"].ChannelName
					ccn = v.UserOperation["GetCompanyInfo"].CCName
					fcn = v.UserOperation["GetCompanyInfo"].Fcn
					break
				}
			}
			fmt.Println("cn,ccn,fcn getcompanyinfo is", cn, ccn, fcn)
			key := webutil.PhoneType + r.FormValue("bnumber")

			companyInfo, err := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
			if err != nil {
				http.Error(w, "Unable to invoke hello in the blockchain", 500)
			}
			json.Unmarshal([]byte(companyInfo), &cInfo)
			fmt.Println("companyInfo is cinfo is", companyInfo, cInfo)
		}
		test := webutil.Person{"john", "boy"}
		data["CompanyInfo"] = cInfo
		data["Why"] = test
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	fmt.Println("org and username", oName, uName)
	batch := app.GetBatchInfo(oName, uName)
	data["BatchInfo"] = batch
	renderTemplate(w, r, "getcompany.html", data)
}

func (app *Application) AddCompanySubcomponentHandler(w http.ResponseWriter, r *http.Request) {

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
					cn = v.UserOperation["AddSupplier"].ChannelName
					ccn = v.UserOperation["AddSupplier"].CCName
					fcn = v.UserOperation["AddSupplier"].Fcn
					break
				}
			}
			key1 := webutil.PhoneType
			key2 := r.FormValue("bnumber")
			key := key1 + key2

			sn := r.FormValue("subcomponetName")
			sl := r.FormValue("companyName")
			ci := r.FormValue("location")
			//add properties to args
			passargs = append(passargs, key)
			passargs = append(passargs, sn)
			passargs = append(passargs, sl)
			passargs = append(passargs, ci)
			fmt.Println("add sub aaaa", sn, sl, ci, cn, ccn, fcn)
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
	renderTemplate(w, r, "addcompanysubcomponent.html", data)
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
