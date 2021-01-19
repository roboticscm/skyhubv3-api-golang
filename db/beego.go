package db

import (
	"backend/system/config"
	. "backend/system/models"
	"backend/system/slog"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var DBConnectionStr string

type BeegoDB struct {
}

func (db *BeegoDB) Init(conf config.CommonConfiguration) {
	orm.RegisterModel(new(LocaleResource))
	orm.RegisterModel(new(RefreshToken))
	orm.RegisterModel(new(Account))
	orm.RegisterModel(new(MenuControl))
	orm.RegisterModel(new(MenuHistory))
	orm.RegisterModel(new(Role))
	orm.RegisterModel(new(SkyLog))
	orm.RegisterModel(new(UserSetting))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	DBConnectionStr = fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable", conf.DBUser, conf.DBPassword, conf.DBServer, conf.DBPort, conf.DBName)
	slog.Compaq(DBConnectionStr)
	orm.RegisterDataBase("default",
		"postgres",
		DBConnectionStr)
}

func (db *BeegoDB) Sync() error {
	name := "default"
	// IMPORTANT: true value will drop your DB
	force := false
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
