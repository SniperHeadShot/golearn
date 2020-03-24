package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ConnectInfo struct {
	Type     *string `json:"type"`
	Host     *string `json:"host"`
	Port     *int    `json:"port"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	DbName   *string `json:"dbName"`
}

// 建立Mysql连接
func BuildMysqlDbConnect(connInfo *ConnectInfo) (*sql.DB, interface{}) {
	if connInfo == nil {
		return nil, "db connect info is empty"
	}
	path := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", *connInfo.Username, *connInfo.Password, *connInfo.Host, *connInfo.Port, *connInfo.DbName)
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// 最大打开连接数
	db.SetMaxOpenConns(20)
	// 最大空闲连接数
	db.SetMaxIdleConns(10)
	// 连接过期时间
	db.SetConnMaxLifetime(time.Second * 10)
	return db, nil
}

// 根据类型获取对应连接
func BuildDbWithType(dbType, dbName *string) (*sql.DB, interface{}) {
	if IsEmpty(dbType) || IsEmpty(dbName) {
		return nil, "dbType or dbName is empty"
	}
	// 读取配置
	var connectInfos []ConnectInfo
	err := ReadConfigFromFile(&connectInfos)
	if err != nil {
		return nil, err
	}
	var connectConfig *ConnectInfo
	// 遍历配置
	for _, connectInfo := range connectInfos {
		if *connectInfo.Type == *dbType {
			connectConfig = &connectInfo
			break
		}
	}
	if connectConfig == nil {
		return nil, fmt.Sprintf("invalid dbType %s!", *dbType)
	}
	connectConfig.DbName = dbName
	return BuildMysqlDbConnect(connectConfig)
}
