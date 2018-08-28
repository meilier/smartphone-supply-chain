package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
)

//HomeHandler : home page
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	var supplierInfo webutil.CompanyInfo
	uName := webutil.MySession.GetUserName(r)
	oName := webutil.MySession.GetOrgName(r)
	if len(uName) == 0 {
		http.Redirect(w, r, "./login.html", 302)
		return
	}
	fmt.Println(len(uName))

	fmt.Println("unameis", "gg", uName, "gg")
	fmt.Printf("byte uname  is %x", uName)
	var cn string
	var ccn string
	var fcn string

	for _, v := range webutil.Orgnization[oName] {
		fmt.Println("org user", v.UserName)
		if v.UserName == uName {
			cn = v.UserOperation["GetSupplier"].ChannelName
			ccn = v.UserOperation["GetSupplier"].CCName
			fcn = v.UserOperation["GetSupplier"].Fcn
			fmt.Println("query channel is ", cn)
			break
		}
	}

	supplierValue, err := app.Fabric[uName].QueryCC(cn, ccn, fcn, []byte("smartisan-u2-pro-zuzhuang"))
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
