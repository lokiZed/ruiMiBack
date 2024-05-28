package database

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var sqlConn sqlx.SqlConn

func GetMysqlConn() sqlx.SqlConn {
	return sqlConn
}

func initMysql(dsn string) error {
	sqlConn = sqlx.NewMysql(dsn)
	return nil
}
