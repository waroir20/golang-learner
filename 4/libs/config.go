package database

import (
	"errors"
	"fmt"
)

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//---------------------------------------------------------------------------------------

// Config - simple model containing the minimum required details to build a connection string
type Config struct {
	Host          string
	Port          int
	EnableSSL     bool
	DatabaseName  string
	DefaultSchema string
	Username      string
	Password      string
}

// ConnectionString - Load database configuration URI
func (d Config) ConnectionString() (string, error) {
	dsn := ""
	if d.Host != "" {
		dsn = fmt.Sprintf("host='%s'", d.Host)
	} else {
		return "", errors.New("no host provided for connection string")
	}
	if d.Port != 0 {
		dsn = fmt.Sprintf("%s port=%d", dsn, d.Port)
	} else {
		dsn = fmt.Sprintf("%s port=5432", dsn)
	}
	if d.DatabaseName != "" {
		dsn = fmt.Sprintf("%s dbname='%s'", dsn, d.DatabaseName)
	}
	if d.Username != "" {
		dsn = fmt.Sprintf("%s user='%s'", dsn, d.Username)
	}
	if d.Password != "" {
		dsn = fmt.Sprintf("%s password='%s'", dsn, d.Password)
	}
	if d.EnableSSL {
		dsn = fmt.Sprintf("%s sslmode='require'", dsn)
	} else {
		dsn = fmt.Sprintf("%s sslmode='disable'", dsn)
	}
	return dsn, nil
}
