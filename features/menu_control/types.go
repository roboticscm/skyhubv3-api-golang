package menu_control

type MenuControl struct {
	ControlId int64   `json:"controlId"`
	Code      *string `json:"code" orm:"null"`
	Name      *string `json:"name" orm:"null"`
	Checked   *bool   `json:"checked" orm:"null"`
}

type Control struct {
	ControlId int64 `json:"controlId"`
	Checked   *bool `json:"checked" orm:"null"`
}

type MenuControlBody struct {
	MenuPath     string    `json:"menuPath"`
	MenuControls []Control `json:"menuControls" orm:"null"`
}
