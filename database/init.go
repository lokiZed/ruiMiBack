package database

import "ruiMiBack2/internal/config"

func InitDatabase(cfg config.DataBase) error {
	if err := initMysql(cfg.Mysql.Dsn); err != nil {
		return err
	}
	return nil
}
