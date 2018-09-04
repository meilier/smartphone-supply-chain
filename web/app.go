package web

import (
	"fmt"
	"net/http"

	"github.com/meilier/smartphone-supply-chain/web/controllers"
)

func Serve(app *controllers.Application) {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//smartphone
	http.HandleFunc("/addbatch.html", app.AddBatchHandler)
	http.HandleFunc("/getbatch.html", app.GetBatchHandler)
	http.HandleFunc("/deletebatch.html", app.DeleteBatchHandler)
	// http.HandleFunc("/getbattery.html", app.AddBatchHandler)
	// http.HandleFunc("/addbatch.html", app.AddBatchHandler)
	// http.HandleFunc("/addbatch.html", app.AddBatchHandler)
	// http.HandleFunc("/addbatch.html", app.AddBatchHandler)
	// http.HandleFunc("/addbatch.html", app.AddBatchHandler)

	http.HandleFunc("/home.html", app.HomeHandler)
	//http.HandleFunc("/addsupplier.html", app.AddSupplierHandler)
	http.HandleFunc("/register.html", app.RegisterHandler)
	http.HandleFunc("/login.html", app.LoginHandler)
	http.HandleFunc("/logout.html", app.LogoutHandler)
	http.HandleFunc("/getbattery.html", app.GetBatteryHandler)
	http.HandleFunc("/getdisplay.html", app.GetDisplayHandler)
	http.HandleFunc("/getcpu.html", app.GetCpuHandler)

	//battery,display,cpu
	http.HandleFunc("/homebattery.html", app.HomeBatteryHandler)
	http.HandleFunc("/addcompany.html", app.AddCompanyHandler)
	http.HandleFunc("/getcompany.html", app.GetCompanyHandler)
	http.HandleFunc("/addcompanysubcomponent.html", app.AddCompanySubcomponentHandler)
	http.HandleFunc("/deletesubcomponent.html", app.DeleteCompanySubcomponentHandler)

	//assembly
	http.HandleFunc("/homeassembly.html", app.HomeAssemblyHandler)
	http.HandleFunc("/getassembly.html", app.GetAssemblyHandler)
	http.HandleFunc("/addassembly.html", app.AddAssemblyHandler)
	http.HandleFunc("/deleteassembly.html", app.DeleteAssemblyHandler)

	//logistics
	http.HandleFunc("/homelogistics.html", app.HomeLogisticsHandler)
	http.HandleFunc("/getlogistics.html", app.GetLogisticsHandler)
	http.HandleFunc("/addlogistics.html", app.AddLogisticsHandler)
	http.HandleFunc("/deletelogistics.html", app.DeleteLogisticsHandler)

	//sales
	http.HandleFunc("/homesales.html", app.HomeSalesHandler)
	http.HandleFunc("/getsales.html", app.GetSalesHandler)
	http.HandleFunc("/addsales.html", app.AddSalesHandler)
	http.HandleFunc("/deletesales.html", app.DeleteSalesHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login.html", http.StatusTemporaryRedirect)
	})

	//admin
	http.HandleFunc("/queryphoneinfo.html", app.QueryPhoneHandler)

	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)

	// open other services at the same time, such as Store listening at 3003
}
