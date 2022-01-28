package dao

import (
	"fmt"
	"github.com/readygo67/LiquidationBot/config"
	"github.com/readygo67/LiquidationBot/dao/model"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql" // nolint
	"gorm.io/gorm"
)

// DB is the global mysql connection instance.
var DB *gorm.DB

// InitDB initializes the DB instance and migrates the table schema.
func InitDB(conf *config.DB) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC", conf.UserName, conf.Password, conf.Host, conf.Port, conf.Database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
		// panic(err)
	}

	// DB = DB.Debug()

	DB.AutoMigrate(&model.Account{}, &model.Prices{}, &model.KV{})
	return nil
}
