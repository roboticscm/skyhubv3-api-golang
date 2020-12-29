package convertor

import (
	"strconv"
)

func StringToInt64(source interface{}) (int64, error) {
	num, err := strconv.ParseInt(source.(string), 10, 64)
	if err != nil {
		return -1, err
	}
	return num, nil
}
