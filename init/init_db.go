package init

import (
	"github.com/shanlongpan/catgin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	var err error
	config.DB, err = gorm.Open(mysql.Open("root:zhaoliqing@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
