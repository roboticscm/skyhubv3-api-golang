package skylog

// l.id, l.created_at as date, a.username as user, l.reason, l.description, l.short_description, '' as view
type SkyLog struct {
	Id               int64   `json:"id"`
	Date             *int64  `json:"date" orm:"null"`
	User             *string `json:"user" orm:"null"`
	Reason           *string `json:"reason" orm:"null"`
	Description      *string `json:"description" orm:"null"`
	ShortDescription *string `json:"shortDescription" orm:"null"`
	View             *string `json:"view" orm:"null"`
}
