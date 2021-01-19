package menu

type Menu struct {
	Id   int64   `json:"id"`
	Code *string `json:"code" orm:"null"`
	Name *string `json:"name" orm:"null"`
	Path *string `json:"path" orm:"null"`
}
