package lib

import (
	"fmt"
	"strconv"
)

func ToInt64(source interface{}) (int64, error) {
	if source == nil {
		return 0, nil
	}

	str := fmt.Sprintf("%v", source)

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return -1, err
	}
	return num, nil
}
