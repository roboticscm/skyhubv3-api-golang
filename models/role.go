package models

type Role struct {
	Id        int64   `json:"id"`
	OrgId     *int64  `json:"orgId" orm:"null"`
	Code      *string `json:"code" orm:"null"`
	Name      *string `json:"name" orm:"null"`
	Sort      *int32  `json:"sort" orm:"null"`
	Disabled  *bool   `json:"disabled" orm:"null"`
	CreatedBy *int64  `json:"createdBy" orm:"null"`
	CreatedAt *int64  `json:"createdAt" orm:"null"`
	UpdatedBy *int64  `json:"updatedBy" orm:"null"`
	UpdatedAt *int64  `json:"updatedAt" orm:"null"`
	DeletedBy *int64  `json:"deletedBy" orm:"null"`

	Version *int32 `json:"version" orm:"null"`
}
