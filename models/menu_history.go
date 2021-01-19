package models

type MenuHistory struct {
	Id         int64  `json:"id"`
	MenuId     *int64 `json:"menuId" orm:"null"`
	DepId      *int64 `json:"depId" orm:"null"`
	AccountId  *int64 `json:"accountId" orm:"null"`
	LastAccess *int64 `json:"lastAccess" orm:"null"`
	Disabled   *bool  `json:"disabled" orm:"null"`
	CreatedBy  *int64 `json:"createdBy" orm:"null"`
	CreatedAt  *int64 `json:"createdAt" orm:"null"`
	UpdatedBy  *int64 `json:"updatedBy" orm:"null"`
	UpdatedAt  *int64 `json:"updatedAt" orm:"null"`
	DeletedBy  *int64 `json:"deletedBy" orm:"null"`
	DeletedAt  *int64 `json:"deletedAt" orm:"null"`
	Version    *int32 `json:"version" orm:"null"`
}
