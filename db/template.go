package db

import (
	"github.com/astaxie/beego/orm"
)

func SelectJson(sql string, param ...interface{}) (interface{}, error) {
	o := orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(sql, param...).Values(&maps)

	if err != nil {
		return nil, err
	}

	return maps[0]["json"], nil
}
