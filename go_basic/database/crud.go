package database

import (
    "context"
    "database/sql"
    "fmt"
    // rand "math/rand/v2"
    // "time"

    _ "github.com/go-sql-driver/mysql" //init()
    // gsb "github.com/huandu/go-sqlbuilder"
)

// insert 插入数据
func Insert(db *sql.DB) {
    //一条sql,插入2行记录
    res, err := db.Exec("insert into student (name,province,city,enrollment) values ('小明', '深圳', '深圳', '2022-07-03'), ('小红', '上海', '上海', '2022-07-03')")
    CheckError(err)
    lastId, err := res.LastInsertId() //ID自增,用过的id(即使对应的行已delete)不会重复使用.如果使用单个INSERT语句将多行插入到表中,则LastInsertId是第一条数据使用的id
    CheckError(err)
    fmt.Printf("after insert last id %d\n", lastId)
    rows, err := res.RowsAffected() //插入2行,所以影响了2行
    CheckError(err)
    fmt.Printf("insert affect %d row\n", rows)
}

// replace 插入(覆盖)数据
func Replace(db *sql.DB) {
    //由于name字段上有唯一索引,insert重复的name会报错.而使用replace会先删除,再插入
    res, err := db.Exec("replace into student (name,province,city,enrollment) values ('小明', '深圳', '深圳', '2025-07-03'), ('小红', '上海', '上海', '2005-07-03')")
    CheckError(err)
    lastId, err := res.LastInsertId() //ID自增,用过的id(即使对应的行已delete)不会重复使用
    CheckError(err)
    fmt.Printf("after insert last id %d\n", lastId)
    rows, err := res.RowsAffected() //先删除,后插入,影响了4行
    CheckError(err)
    fmt.Printf("insert affect %d row\n", rows)
}

// update 修改数据
func Update(db *sql.DB) {
    //不同的city加不同的分数
    res, err := db.Exec("update student set score=score+10 where city='上海'") //上海加10分
    CheckError(err)
    lastId, err := res.LastInsertId() //0,仅插入操作才会给LastInsertId赋值
    CheckError(err)
    fmt.Printf("after update last id %d\n", lastId)
    rows, err := res.RowsAffected() //where city=?命中了几行,就会影响几行
    CheckError(err)
    fmt.Printf("update affect %d row\n", rows)
}

// 事务
func Transaction(db *sql.DB) {
    tx, err := db.BeginTx(context.Background(), nil) // 开始事务
    CheckError(err)
    _, err = tx.Exec("insert into student (name,province,city,enrollment,score) values ('Tom', '深圳', '深圳', '2022-07-03', 40)")
    CheckError(err)
    // _, err = tx.Exec("insert into student (name,province,city,enrollment,score) values ('Tom', '深圳', '深圳', '2022-07-03', 40)") // 一旦中间某一步出错失败,则事务里的所有操作全部回滚
    // CheckError(err)
    if err = tx.Commit(); err != nil { //整体提交
        fmt.Println("第一次commit失败", err)
    }

    tx.Exec("insert into student (name,province,city,enrollment,score) values ('Lily', '深圳', '深圳', '2022-07-03', 40)")
    if err = tx.Commit(); err != nil { //整体提交
        fmt.Println("第二次commit失败", err) //commit或Rollback只能执行一次
    }
}

// delete 删除数据
func Delete(db *sql.DB) {
    res, err := db.Exec("delete from student where id=6") //删除id为6的记录
    CheckError(err)
    rows, err := res.RowsAffected() //where id=6命中了几行,就会影响几行
    CheckError(err)
    fmt.Printf("delete affect %d row(s)\n", rows)
}

type User struct {
    Id      int
    Gender  string
    Score   float64
}
