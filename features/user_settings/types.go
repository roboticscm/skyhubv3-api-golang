package user_settings

type UserSetting struct {
	CompanyId    *int64  `json:"companyId"`
	CompanyName  *string `json:"companyName" orm:"null"`
	BranchId     *int64  `json:"branchId" orm:"null"`
	BranchName   *string `json:"branchName" orm:"null"`
	Locale       *string `json:"locale" orm:"null"`
	Theme        *string `json:"theme" orm:"null"`
	DepartmentId *string `json:"departmentId" orm:"null"`
	MenuPath     *string `json:"menuPath" orm:"null"`
}

type UserSettingBody struct {
	BranchId  interface{} `json:"branchId" orm:"null"`
	MenuPath  *string     `json:"menuPath" orm:"null"`
	ElementId *string     `json:"elementId" orm:"null"`
	Keys      []string    `json:"keys" orm:"null"`
	Values    []string    `json:"values" orm:"null"`
}
