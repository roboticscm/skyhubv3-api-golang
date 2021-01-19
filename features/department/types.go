package department

type Department struct {
	Id   int64   `json:"id"`
	Name *string `json:"name" orm:"null"`
}

type DepartmentId struct {
	DepId *int64 `json:"depId"`
}
