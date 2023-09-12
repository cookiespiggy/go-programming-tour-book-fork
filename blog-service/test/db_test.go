package test

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestCreateDB(t *testing.T) {
	// 连接到MySQL数据库
	db, err := sql.Open("mysql", "root:pwd@tcp(ip:port)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// 创建名为blog_service的数据库，并使用utf8编码
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS blog_service DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci")
	if err != nil {
		panic(err.Error())
	}

	// 输出成功信息
	fmt.Println("blog_service数据库创建成功")
}
