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

func init() {
	Orgnization = make(map[string][]OrgMember)

	wzxop := make(map[string]BaseFuncInfo)
	wzxop["AddSupplier"] = BaseFuncInfo{"supplychannel", "addsupplier", "addSupplier"}
	wzxop["GetSupplier"] = BaseFuncInfo{"supplychannel", "addsupplier", "getSupplier"}
	wzxChannel := []string{"supplychannel", "assemblychannel", "logisticschannel", "salechannel"}
	mwzx := OrgMember{"wzx", "arclabw401wzx", wzxChannel, wzxop, "./profile/smartphone/connection-profile-wzx.yaml"}
	Orgnization["smartphone"] = append(Orgnization["smartphone"], mwzx)

	lwhop := make(map[string]BaseFuncInfo)
	lwhop["AddSupplier"] = BaseFuncInfo{"supplychannel", "ccbattery", "record"}
	lwhop["AddSupplier"] = BaseFuncInfo{"supplychannel", "ccbattery", "query"}
	lwhChannel := []string{"supplierchannel"}
	mlwh := OrgMember{"lwh", "arclabw401lwh", lwhChannel, wzxop, "./profile/supplier/connection-profile-lwh.yaml"}
	Orgnization["supplier"] = append(Orgnization["supplier"], mlwh)

	wyhop := make(map[string]BaseFuncInfo)
	wyhop["AddSupplier"] = BaseFuncInfo{"supplychannel", "addsupplier", "addSupplier"}
	wyhop["AddSupplier"] = BaseFuncInfo{"supplychannel", "addsupplier", "getSupplier"}
	wyhChannel := []string{"assemblychannel"}
	mwyh := OrgMember{"wyh", "arclabw401wyh", wyhChannel, wyhop, "./profile/assembly/connection-profile-wyh.yaml"}
	Orgnization["assembly"] = append(Orgnization["assembly"], mwyh)

	yzxop := make(map[string]BaseFuncInfo)
	yzxop["AddSupplier"] = BaseFuncInfo{"supplychannel", "addsupplier", "addSupplier"}
	yzxop["AddSupplier"] = BaseFuncInfo{"supplychannel", "addsupplier", "getSupplier"}
	yzxChannel := []string{"logisticschannel"}
	myzx := OrgMember{"yzx", "arclabw401yzx", yzxChannel, yzxop, "./profile/logistics/connection-profile-yzx.yaml"}
	Orgnization["logistics"] = append(Orgnization["logistics"], myzx)

	xjxop := make(map[string]BaseFuncInfo)
	xjxop["AddSupplier"] = BaseFuncInfo{"supplychannel", "addsupplier", "addSupplier"}
	xjxop["AddSupplier"] = BaseFuncInfo{"supplychannel", "addsupplier", "getSupplier"}
	xjxChannel := []string{"supplierchannel"}
	mxjx := OrgMember{"xjx", "arclabw401xjx", xjxChannel, xjxop, "./profile/store/connection-profile-xjx.yaml"}
	Orgnization["store"] = append(Orgnization["logistics"], mxjx)

}
