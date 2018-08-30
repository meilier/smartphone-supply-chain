package webutil

//CompanyInfo define the company structure, with x properties.  Structure tags are used by encoding/json library
type CompanyInfo struct {
	ConcreteCompanyInfo []BaseCompanyInfo `json:"concretecompanyinfo"`
}

//BaseCompanyInfo define basic company info
type BaseCompanyInfo struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type AssemblyInfo struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Manager  string `json:"manager"`
	Date     string `json:"date"`
}

var AphoneSerialNumber = [5]string{"10000000", "10000001", "10000002", "10000003", "10000004"}
var XphoneSerialNumber = [5]string{"20000000", "20000001", "20000002", "20000003", "20000004"}

type BaseFuncInfo struct {
	ChannelName string
	CCName      string
	Fcn         string
}

type OrgOperation map[string]BaseFuncInfo

type OrgMember struct {
	UserName      string
	Secret        string
	ChannelName   []string
	UserOperation OrgOperation
	FilePath      string
}

var Orgnization map[string][]OrgMember
var PhoneType = "Aphone"

func init() {
	Orgnization = make(map[string][]OrgMember)
	//-----------smartphone-------------
	wzxop := make(map[string]BaseFuncInfo)

	//for batch
	wzxop["AddBatchBattery"] = BaseFuncInfo{"batterychannel", "addbattery", "addBatchInfo"}
	wzxop["GetBatchBattery"] = BaseFuncInfo{"batterychannel", "addbattery", "getBatchInfo"}

	wzxop["AddBatchDisplay"] = BaseFuncInfo{"displaychannel", "adddisplay", "addBatchInfo"}
	wzxop["GetBatchDisplay"] = BaseFuncInfo{"displaychannel", "adddisplay", "getBatchInfo"}

	wzxop["AddBatchCpu"] = BaseFuncInfo{"cpuchannel", "addcpu", "addBatchInfo"}
	wzxop["GetBatchCpu"] = BaseFuncInfo{"cpuchannel", "addcpu", "getBatchInfo"}

	wzxop["AddBatchAssembly"] = BaseFuncInfo{"assemblychannel", "addassembly", "addBatchInfo"}
	wzxop["GetBatchAssembly"] = BaseFuncInfo{"assemblychannel", "addassembly", "getBatchInfo"}

	wzxop["AddBatchLogistics"] = BaseFuncInfo{"logisticschannel", "addlogistics", "addBatchInfo"}
	wzxop["GetBatchLogistics"] = BaseFuncInfo{"logisticschannel", "addlogistics", "getBatchInfo"}

	//for supply chain info
	wzxop["GetBatterytInfo"] = BaseFuncInfo{"batterychannel", "addbattery", "getSupplier"}
	wzxop["GetDisplayInfo"] = BaseFuncInfo{"displaychannel", "adddisplay", "getSupplier"}
	wzxop["GetCpuInfo"] = BaseFuncInfo{"cpuchannel", "addcpu", "getSupplier"}
	wzxop["GetAssemblyInfo"] = BaseFuncInfo{"assemblychannel", "addassembly", "getSupplier"}
	wzxop["GetLogisticsInfo"] = BaseFuncInfo{"logisticschannel", "addlogistics", "getSupplier"}
	wzxop["GetStore"] = BaseFuncInfo{"saleschannel", "addsale", "getSupplier"}

	wzxChannel := []string{"batterychannel", "displaychannel", "cpuchannel", "assemblychannel", "logisticschannel", "salechannel"}
	mwzx := OrgMember{"wzx", "arclabw401wzx", wzxChannel, wzxop, "./profile/smartphone/connection-profile-wzx.yaml"}
	Orgnization["smartphone"] = append(Orgnization["smartphone"], mwzx)

	//--------------battery---------------
	lwhop := make(map[string]BaseFuncInfo)
	lwhop["AddSupplier"] = BaseFuncInfo{"batterychannel", "addbattery", "addSupplier"}
	lwhop["GetSupplier"] = BaseFuncInfo{"batterychannel", "addbattery", "getSupplier"}
	lwhop["AddCompanyInfo"] = BaseFuncInfo{"batterychannel", "addbattery", "addCompanyInfo"}
	lwhop["GetCompanyInfo"] = BaseFuncInfo{"batterychannel", "addbattery", "getCompanyInfo"}
	lwhChannel := []string{"batterychannel"}
	mlwh := OrgMember{"lwh", "arclabw401lwh", lwhChannel, lwhop, "./profile/battery/connection-profile-lwh.yaml"}
	Orgnization["battery"] = append(Orgnization["battery"], mlwh)

	//-------------display-----------------
	lwhop1 := make(map[string]BaseFuncInfo)
	lwhop1["AddSupplier"] = BaseFuncInfo{"displaychannel", "adddisplay", "addSupplier"}
	lwhop1["GetSupplier"] = BaseFuncInfo{"displaychannel", "adddisplay", "getSupplier"}
	lwhop1["AddCompanyInfo"] = BaseFuncInfo{"displaychannel", "adddisplay", "addCompanyInfo"}
	lwhop1["GetCompanyInfo"] = BaseFuncInfo{"displaychannel", "adddisplay", "getCompanyInfo"}
	lwhChannel1 := []string{"displaychannel"}
	mlwh1 := OrgMember{"lwh", "arclabw401lwh", lwhChannel1, lwhop1, "./profile/display/connection-profile-lwh.yaml"}
	Orgnization["display"] = append(Orgnization["display"], mlwh1)

	//--------------cpu------------------
	lwhop2 := make(map[string]BaseFuncInfo)
	lwhop2["AddSupplier"] = BaseFuncInfo{"cpuchannel", "addcpu", "addSupplier"}
	lwhop2["GetSupplier"] = BaseFuncInfo{"cpuchannel", "addcpu", "getSupplier"}
	lwhop2["AddCompanyInfo"] = BaseFuncInfo{"cpuchannel", "addcpu", "addCompanyInfo"}
	lwhop2["GetCompanyInfo"] = BaseFuncInfo{"cpuchannel", "addcpu", "getCompanyInfo"}
	lwhChannel2 := []string{"cpuchannel"}
	mlwh2 := OrgMember{"lwh", "arclabw401lwh", lwhChannel2, lwhop2, "./profile/cpu/connection-profile-lwh.yaml"}
	Orgnization["cpu"] = append(Orgnization["cpu"], mlwh2)

	//-------------assembly---------------
	wyhop := make(map[string]BaseFuncInfo)
	wyhop["AddAssembly"] = BaseFuncInfo{"assemblychannel", "addassembly", "addAssembly"}
	wyhop["GetAssembly"] = BaseFuncInfo{"assemblychannel", "addassembly", "getAssembly"}
	wyhChannel := []string{"assemblychannel"}
	mwyh := OrgMember{"wyh", "arclabw401wyh", wyhChannel, wyhop, "./profile/assembly/connection-profile-wyh.yaml"}
	Orgnization["assembly"] = append(Orgnization["assembly"], mwyh)

	//-------------logistics---------------
	yzxop := make(map[string]BaseFuncInfo)
	yzxop["AddLogistics"] = BaseFuncInfo{"logisticschannel", "addlogistics", "addLogistics"}
	yzxop["GetLogistics"] = BaseFuncInfo{"logisticschannel", "addlogistics", "getLogistics"}
	yzxChannel := []string{"logisticschannel"}
	myzx := OrgMember{"yzx", "arclabw401yzx", yzxChannel, yzxop, "./profile/logistics/connection-profile-yzx.yaml"}
	Orgnization["logistics"] = append(Orgnization["logistics"], myzx)

	//---------------sales---------------
	xjxop := make(map[string]BaseFuncInfo)
	xjxop["AddSale"] = BaseFuncInfo{"saleschannel", "addsale", "addSale"}
	xjxop["AddSale"] = BaseFuncInfo{"saleschannel", "addsale", "getSale"}
	xjxChannel := []string{"saleschannel"}
	mxjx := OrgMember{"xjx", "arclabw401xjx", xjxChannel, xjxop, "./profile/sales/connection-profile-xjx.yaml"}
	Orgnization["sales"] = append(Orgnization["sales"], mxjx)

}
