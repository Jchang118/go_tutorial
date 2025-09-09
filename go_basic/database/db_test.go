package database_test

import (
    "database/sql"
    "go_tutorial/go_basic/database"
    // "fmt"
    "testing"
)

var (
    db *sql.DB
)

func init() {
    var err error
    /*
    DSN(data source name)格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
    例如user:password@tcp(localhost:5555)/dbname?charset=utf8
    如果是本地MySQL,且采用默认的3306端口,可简写为: user:password@/dbname
    想要正确的处理time.Time,你需要带上parseTime参数
    要支持完整的UTF-8编码,你需要将charset=utf8更改为charset=utf8mb4
    loc=Local采用机器本地的时区
    */
    // db, err := sql.Open("mysql", "tester:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
    // db可以并发使用
    db, err = sql.Open("mysql", "tester:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Australia%2FSydney")
    database.CheckError(err)
    // defer db.Close()
}

func TestInsert(t *testing.T) {
    database.Insert(db)
}

func TestReplace(t *testing.T) {
    database.Replace(db)
}

func TestUpdate(t *testing.T) {
    database.Update(db)
}

func TestDelete(t *testing.T) {
    database.Delete(db)
}

func TestTransaction(t *testing.T) {
    database.Transaction(db)
}
