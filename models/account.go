package models

type Account struct {
	Id        int64   `json:"id"`
	Username    *string `json:"username" orm:"null"`
	Password    *string `json:"password" orm:"null"`
	Disabled  *bool   `json:"disabled" orm:"null"`
	CreatedBy *int64  `json:"createdBy" orm:"null"`
	CreatedAt *int64  `json:"createdAt" orm:"null"`
	UpdatedBy *int64  `json:"updatedBy" orm:"null"`
	UpdatedAt *int64  `json:"updatedAt" orm:"null"`
	DeletedBy *int64  `json:"deletedBy" orm:"null"`

	Version *int32 `json:"version" orm:"null"`
}
