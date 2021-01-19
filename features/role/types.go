package role

type RoleBody struct {
	Id    *int64      `json:"id"`
	Code  *string     `json:"code"`
	Name  *string     `json:"name"`
	Sort  *int32      `json:"sort"`
	OrgId interface{} `json:"orgId"`
}
