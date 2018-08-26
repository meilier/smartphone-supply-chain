package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/meilier/smartphone-supply-chain/web/webutil"

	"github.com/meilier/smartphone-supply-chain/blockchain"
)

type Application struct {
	Fabric *blockchain.FabricSetup
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	lp := filepath.Join("web", "templates", "layout.html")
	tp := filepath.Join("web", "templates", templateName)

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
	fmt.Println("first")
	resultTemplate, err := template.ParseFiles(tp, lp)
	if err != nil {
		// Log the detailed error
		fmt.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	newData := data.(map[string]interface{})

	//get session
	uname := webutil.MySession.GetUserName(r)
	uorg := webutil.MySession.GetOrgName(r)
	fmt.Println("uname is ", uname)
	if uname != "" {
		//set data
		newData["Username"] = uname
		newData["LoginStatus"] = true
		newData["OrgName"] = uorg
	} else {
		//redirect to login
		http.Redirect(w, r, "./login.html", 302)
		newData["Username"] = ""
		newData["OrgName"] = ""
		newData["LoginStatus"] = false
	}
	fmt.Println("second")
	if err := resultTemplate.ExecuteTemplate(w, "layout", newData); err != nil {
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
