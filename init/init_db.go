package init

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"fmt"
	"github.com/shanlongpan/catgin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	var err error
	//"root:xxx@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	// 多库需要配置多个连接
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Conf.Mysql.Username, config.Conf.Mysql.Password, config.Conf.Mysql.Dsn, config.Conf.Mysql.Port, config.Conf.Mysql.Database)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
