package pg

import "gorm.io/gorm"

type ConnectionConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
	TimeZone string
}

type GormConfig = gorm.Config

var defaultConfig = ConnectionConfig{
	Host:     "localhost",
	User:     "postgres",
	Password: "postgres",
	DBName:   "postgres",
	Port:     5432,
	SSLMode:  "disable",
	TimeZone: "UTC",
}
