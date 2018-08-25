package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/meilier/smartphone-supply-chain/web/webutil"
	"github.com/mitchellh/mapstructure"

	"github.com/meilier/smartphone-supply-chain/blockchain"
)

type Application struct {
	Fabric *blockchain.FabricSetup
}

type BasicData struct {
	Username     string
	LoginStatus  bool
	SupplierInfo string
}

type SupplierData struct {
	Username     string
	LoginStatus  bool
	SupplierInfo string
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	lp := filepath.Join("web", "templates", "layout.html")
	tp := filepath.Join("web", "templates", templateName)

	// factory create struct
	mData := new(SupplierData)
	// Return a 404 if the template doesn't exist
	info, err := os.Stat(tp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	resultTemplate, err := template.ParseFiles(tp, lp)
	if err != nil {
		// Log the detailed error
		fmt.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	//get session
	uname := webutil.MySession.GetUserName(r)
	var mMap map[string]interface{}
	mMap = make(map[string]interface{})
	if uname != "" {
		//set data
		mData.Username = uname
		mData.LoginStatus = true
	} else {
		//redirect to login
		mData.Username = ""
		mData.LoginStatus = false
	}
	//获取interface中的数据
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		mMap[t.Field(i).Name] = v.Field(i).Interface()
		println(t.Field(i).Name, v.Field(i).Interface())
	}
	//map to struct
	if err := mapstructure.Decode(mMap, &mData); err != nil {
		fmt.Println(err)
	}
	if err := resultTemplate.ExecuteTemplate(w, "layout", mData); err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}

func loginTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	lp := filepath.Join("web", "templates", "layout.html")
	tp := filepath.Join("web", "templates", templateName)
	resultTemplate, _ := template.ParseFiles(tp, lp)
	if err := resultTemplate.ExecuteTemplate(w, "layout", data); err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
