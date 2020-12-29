package models

type LocaleResource struct {
	Id        int64   `json:"id"`
	CompanyId *int64  `json:"companyId" orm:"null"`
	Locale    *string `json:"locale" orm:"null"`
	Category  *string `json:"category" orm:"null"`
	TypeGroup *string `json:"typeGroup" orm:"null"`
	Key       *string `json:"key" orm:"null"`
	Value     *string `json:"value" orm:"null"`
	Sort      *int32  `json:"sort" orm:"null"`
	Disabled  *bool   `json:"disabled" orm:"null"`
	CreatedBy *int64  `json:"createdBy" orm:"null"`
	CreatedAt *int64  `json:"createdAt" orm:"null"`
	UpdatedBy *int64  `json:"updatedBy" orm:"null"`
	UpdatedAt *int64  `json:"updatedAt" orm:"null"`
	DeletedBy *int64  `json:"deletedBy" orm:"null"`

	Version *int32 `json:"version" orm:"null"`
}
