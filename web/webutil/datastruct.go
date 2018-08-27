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
