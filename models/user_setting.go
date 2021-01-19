package models

type UserSetting struct {
	Id        int64   `json:"id"`
	BranchId  *int64  `json:"branchId" orm:"null"`
	AccountId *int64  `json:"accountId" orm:"null"`
	MenuPath  *string `json:"menuPath" orm:"null"`
	ElementId *string `json:"elementId" orm:"null"`
	Key       *string `json:"key" orm:"null"`
	Value     *string `json:"value" orm:"null"`
}
