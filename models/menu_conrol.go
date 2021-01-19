package models

type MenuControl struct {
	Id        int64  `json:"id"`
	MenuId    *int64 `json:"menuId" orm:"null"`
	ControlId *int64 `json:"controlId" orm:"null"`
	Disabled  *bool  `json:"disabled" orm:"null"`
	CreatedBy *int64 `json:"createdBy" orm:"null"`
	CreatedAt *int64 `json:"createdAt" orm:"null"`
	UpdatedBy *int64 `json:"updatedBy" orm:"null"`
	UpdatedAt *int64 `json:"updatedAt" orm:"null"`
	DeletedBy *int64 `json:"deletedBy" orm:"null"`
	DeletedAt *int64 `json:"deletedAt" orm:"null"`
	Version   *int32 `json:"version" orm:"null"`
}
