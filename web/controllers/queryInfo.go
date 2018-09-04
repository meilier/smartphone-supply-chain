package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

func (app *Application) GetBatteryHandler(w http.ResponseWriter, r *http.Request) {

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
		cInfo = app.GetPhoneSupplierInfo(oName, uName, "GetBatteryInfo", r.FormValue("bnumber"))
		data["CompanyInfo"] = cInfo
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	fmt.Println("org and username", oName, uName)
	batch := app.GetPhoneBatchInfo("battery", uName)
	data["BatchInfo"] = batch
	data["BatteryQuery"] = "Query Battery Info"
	renderTemplate(w, r, "getcompany.html", data)
}

func (app *Application) GetDisplayHandler(w http.ResponseWriter, r *http.Request) {

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
		cInfo = app.GetPhoneSupplierInfo(oName, uName, "GetDisplayInfo", r.FormValue("bnumber"))
		data["CompanyInfo"] = cInfo
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	fmt.Println("org and username", oName, uName)
	batch := app.GetPhoneBatchInfo("display", uName)
	data["BatchInfo"] = batch
	data["DisplayQuery"] = "Query Display Info"
	renderTemplate(w, r, "getcompany.html", data)
}

func (app *Application) GetCpuHandler(w http.ResponseWriter, r *http.Request) {

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
		cInfo = app.GetPhoneSupplierInfo(oName, uName, "GetCpuInfo", r.FormValue("bnumber"))
		data["CompanyInfo"] = cInfo
		// txid, err := app.Fabric.InvokeSupplier(passargs)
	}
	fmt.Println("org and username", oName, uName)
	batch := app.GetPhoneBatchInfo("cpu", uName)
	data["BatchInfo"] = batch
	data["CpuQuery"] = "Query Cpu Info"
	renderTemplate(w, r, "getcompany.html", data)
}

func (app *Application) GetPhoneBatchInfo(orgtype string, username string) []string {
	var bInfo webutil.BatchInfo
	//befor send request we need to check session
	if fSetup, ok := app.Fabric[username]; ok {
		var cn string
		var ccn string
		var fcn string
		fmt.Println("username is", username)

		//find cfg name
		//according supplier type to choose corresponding channel
		for _, v := range webutil.Orgnization["smartphone"] {
			fmt.Println("v.Username is", v.UserName)
			if v.UserName == username {
				fmt.Println("aaaaaa suppliertype is ", orgtype)
				switch orgtype {
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

func (app *Application) GetPhoneSupplierInfo(oName, uName, operation, key string) webutil.CompanyInfo {
	var cInfo webutil.CompanyInfo

	if fSetup, ok := app.Fabric[uName]; ok {
		//befor send request we need to check session
		fmt.Println("uname oname haha is", uName, oName)
		var cn string
		var ccn string
		var fcn string
		//find cfg name
		for _, v := range webutil.Orgnization[oName] {
			fmt.Println("v.Username haha is", v.UserName)
			if v.UserName == uName {
				cn = v.UserOperation[operation].ChannelName
				ccn = v.UserOperation[operation].CCName
				fcn = v.UserOperation[operation].Fcn
				break
			}
		}
		fmt.Println("cn,ccn,fcn getcompanyinfo is", cn, ccn, fcn)
		key := webutil.PhoneType + key
		fmt.Println("haha key si", key)

		companyInfo, _ := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
		json.Unmarshal([]byte(companyInfo), &cInfo)
		fmt.Println("companyInfo is", companyInfo)
		fmt.Println("wzx cinfo is", cInfo)
	}
	// txid, err := app.Fabric.InvokeSupplier(passargs)

	fmt.Println("org and username", oName, uName)
	return cInfo
}

func (app *Application) GetPhoneAssemblyInfo(oName, uName, key string) webutil.AssemblyInfo {
	var cInfo webutil.AssemblyInfo

	if fSetup, ok := app.Fabric[uName]; ok {
		//befor send request we need to check session
		fmt.Println("uname oname haha is", uName, oName)
		var cn string
		var ccn string
		var fcn string
		//find cfg name
		for _, v := range webutil.Orgnization[oName] {
			fmt.Println("v.Username haha is", v.UserName)
			if v.UserName == uName {
				cn = v.UserOperation["GetAssembly"].ChannelName
				ccn = v.UserOperation["GetAssembly"].CCName
				fcn = v.UserOperation["GetAssembly"].Fcn
				break
			}
		}
		fmt.Println("cn,ccn,fcn getcompanyinfo is", cn, ccn, fcn)
		key := webutil.PhoneType + key
		fmt.Println("haha key si", key)

		companyInfo, _ := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
		json.Unmarshal([]byte(companyInfo), &cInfo)
		fmt.Println("companyInfo is cinfo is", companyInfo, cInfo)
	}
	// txid, err := app.Fabric.InvokeSupplier(passargs)

	fmt.Println("org and username", oName, uName)
	return cInfo
}

func (app *Application) GetPhoneLogisticsInfo(oName, uName, key string) webutil.TransitInfo {
	var cInfo webutil.TransitInfo

	if fSetup, ok := app.Fabric[uName]; ok {
		//befor send request we need to check session
		fmt.Println("uname oname haha is", uName, oName)
		var cn string
		var ccn string
		var fcn string
		//find cfg name
		for _, v := range webutil.Orgnization[oName] {
			fmt.Println("v.Username haha is", v.UserName)
			if v.UserName == uName {
				cn = v.UserOperation["GetLogistics"].ChannelName
				ccn = v.UserOperation["GetLogistics"].CCName
				fcn = v.UserOperation["GetLogistics"].Fcn
				break
			}
		}
		fmt.Println("cn,ccn,fcn getcompanyinfo is", cn, ccn, fcn)
		key := webutil.PhoneType + key
		fmt.Println("haha key si", key)

		companyInfo, _ := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
		json.Unmarshal([]byte(companyInfo), &cInfo)
		fmt.Println("companyInfo is cinfo is", companyInfo, cInfo)
	}
	// txid, err := app.Fabric.InvokeSupplier(passargs)

	fmt.Println("org and username", oName, uName)
	return cInfo
}

func (app *Application) GetPhoneSalesInfo(oName, uName, key string) webutil.SalesInfo {
	var cInfo webutil.SalesInfo

	if fSetup, ok := app.Fabric[uName]; ok {
		//befor send request we need to check session
		fmt.Println("uname oname haha is", uName, oName)
		var cn string
		var ccn string
		var fcn string
		//find cfg name
		for _, v := range webutil.Orgnization[oName] {
			fmt.Println("v.Username haha is", v.UserName)
			if v.UserName == uName {
				cn = v.UserOperation["GetSales"].ChannelName
				ccn = v.UserOperation["GetSales"].CCName
				fcn = v.UserOperation["GetSales"].Fcn
				break
			}
		}
		fmt.Println("cn,ccn,fcn getcompanyinfo is", cn, ccn, fcn)
		key := webutil.PhoneType + key
		fmt.Println("haha key si", key)

		companyInfo, _ := fSetup.QueryCC(cn, ccn, fcn, []byte(key))
		json.Unmarshal([]byte(companyInfo), &cInfo)
		fmt.Println("companyInfo is cinfo is", companyInfo, cInfo)
	}
	// txid, err := app.Fabric.InvokeSupplier(passargs)

	fmt.Println("org and username", oName, uName)
	return cInfo
}
