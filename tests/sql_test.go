package tests

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

//使用原生库 crud
func Test_ping(t *testing.T) {
	d, _ := sql.Open("mysql", "gosql:gosql@tcp(172.16.16.17:3306)/gosql")
	fmt.Printf("d.Ping(): %v\n", d.Ping())
	//测试建表
	s := `
	CREATE TABLE IF NOT EXISTS gosql.tb  (
		id bigint(20)  NOT NULL,
		name varchar(10),
		PRIMARY KEY (id)
	);
	`
	r, err := d.Exec(s)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("r: %v\n", r)

}
func Test_create_table(t *testing.T) {

}
