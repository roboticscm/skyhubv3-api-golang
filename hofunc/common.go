package hofunc

func Filter(filterFunc func(interface{}) bool, source ...interface{}) []interface{} {
	var dest []interface{}
	for _, v := range source {
		if filterFunc(v) == true {
			dest = append(dest, v)
		}
	}
	return dest
}
