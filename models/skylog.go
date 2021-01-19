package models

type SkyLog struct {
	Id               int64   `json:"id"`
	CompanyId        *int64  `json:"companyId" orm:"null"`
	BranchId         *int64  `json:"branchId" orm:"null"`
	MenuPath         *string `json:"menuPath" orm:"null"`
	IpClient         *string `json:"ipClient" orm:"null"`
	Device           *string `json:"device" orm:"null"`
	Os               *string `json:"os" orm:"null"`
	Browser          *string `json:"browser" orm:"null"`
	ShortDescription *string `json:"shortDescription" orm:"null"`
	Description      *string `json:"description" orm:"null"`
	Reason           *string `json:"reason" orm:"null"`
	CreatedBy        *int64  `json:"createdBy" orm:"null"`
	CreatedAt        *int64  `json:"createdAt" orm:"null"`
	UpdatedBy        *int64  `json:"updatedBy" orm:"null"`
	UpdatedAt        *int64  `json:"updatedAt" orm:"null"`
	DeletedBy        *int64  `json:"deletedBy" orm:"null"`
	DeletedAt        *int64  `json:"deletedAt" orm:"null"`
	Disabled         *bool   `json:"disabled" orm:"null"`
	Version          *int32  `json:"version" orm:"null"`
}
