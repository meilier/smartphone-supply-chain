package controllers

import (
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//QueryPhoneHandler : home page
func (app *Application) QueryPhoneHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	data = make(map[string]interface{})
	var keymap map[string]string
	keymap = make(map[string]string)
	var batch []string
	oName := "smartphone"
	uName := "wzx"
	var binfo webutil.CompanyInfo
	var dinfo webutil.CompanyInfo
	var cinfo webutil.CompanyInfo
	var asinfo webutil.AssemblyInfo
	var trinfo webutil.TransitInfo
	var saInfo webutil.SalesInfo
	if r.FormValue("submitted") == "true" {
		key := r.FormValue("snumber")
		//according snumber to find a batch key
		for k, _ := range webutil.Orgnization {
			if k == "smartphone" {
				continue
			}
			batch = app.GetPhoneBatchInfo(k, uName)
			for _, vb := range batch {
				if key > vb {
					continue
				} else {
					keymap[k] = vb
				}
			}
		}

		binfo = app.GetPhoneSupplierInfo(oName, uName, "GetBatteryInfo", keymap["battery"])
		dinfo = app.GetPhoneSupplierInfo(oName, uName, "GetDisplayInfo", keymap["display"])
		cinfo = app.GetPhoneSupplierInfo(oName, uName, "GetCpuInfo", keymap["cpu"])
		asinfo = app.GetPhoneAssemblyInfo(oName, uName, keymap["assembly"])
		trinfo = app.GetPhoneLogisticsInfo(oName, uName, keymap["logistics"])
		saInfo = app.GetPhoneSalesInfo(oName, uName, key)

		data["BatteryInfo"] = binfo
		data["DisplayInfo"] = dinfo
		data["CpuInfo"] = cinfo
		data["AssemblyInfo"] = asinfo
		data["LogisticsInfo"] = trinfo.ConcreteTransitInfo
		data["SalesInfo"] = saInfo
		fmt.Println("all info is", binfo, dinfo, cinfo, asinfo, trinfo, saInfo)
	}
	queryphoneTemplate(w, r, "queryphoneinfo.html", data)
}
