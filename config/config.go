package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error {
	logger := GetLogger("init")
	var err error

	db, err = InitializeSQLite()
	if err != nil {
		logger.Errorf("sqlite init error: %v", err)
		return err
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
