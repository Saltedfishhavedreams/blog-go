package mysql

import (
	"blog/config"
	"blog/pkg/logger"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var DB *sqlx.DB

func Init(mysqlConfig *config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DbName,
	)

	DB, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		logger.Error("mysql connect error", zap.Error(err))
		return
	}

	DB.SetMaxOpenConns(int(mysqlConfig.MaxOpenConns))
	DB.SetMaxIdleConns(int(mysqlConfig.MaxIdleConns))

	return
}

func Close() {
	DB.Close()
}
