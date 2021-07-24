package db

import (
	"github.com/dsk52/vote/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Port       string
	Connection *gorm.DB
}

type dBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
}

func NewConfig() *dBConfig {
	c := new(dBConfig)
	c.Host = helper.GetEnv("DB_HOST", "127.0.0.1")
	c.Username = helper.GetEnv("DB_USER_NAME", "mysql")
	c.Password = helper.GetEnv("DB_PASSWORD", "mysql")
	c.DBName = helper.GetEnv("DB_PASSWORD", "mysql")
	c.Port = helper.GetEnv("DB_PORT", "3306")
	return c
}

func NewDD() *DB {
	c := NewConfig()

	return newDB(&DB{
		Host:     c.Host,
		Username: c.Username,
		Password: c.Password,
		DBName:   c.DBName,
		Port:     c.Port,
	})
}

func newDB(d *DB) *DB {
	dsn := d.Username + ":" + d.Password + "@tcp(mysql:" + d.Port + ")/" + d.DBName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db

	return d
}
