package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "time"
)
type user struct {
    ID int
    Name string
    Age int
}
// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
    // DSN:Data Source Name
    //dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", "root", "root", "127.0.0.1", "3306", "test", "utf8")
    dsn := "root:root@tcp(127.0.0.1:3306)/test"
    // 不会校验账号密码是否正确
    // 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }

    /*
    最大连接数
    如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。
    如果n<=0，不会限制最大开启连接数，默认为0（无限制）。
     */
    db.SetMaxOpenConns(100)
    /*
    连接池中的最大闲置连接数
    如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。
     */
    db.SetMaxIdleConns(20)
    /*
    最大连接周期
     */
    db.SetConnMaxLifetime(100*time.Second)

    // 尝试与数据库建立连接（校验dsn是否正确）
    err = db.Ping()
    if err != nil {
        log.Fatalln("[db.Ping.Err] ",err)
        return err
    }
    fmt.Println("[db.Ping]","ok")
    return nil
}

// 单行查询
// 单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。
func queryRowDemo() {
    sqlStr := "select id, name, age from user where id=?"
    var u user
    // 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
    err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
    if err != nil {
        fmt.Printf("scan failed, err:%v\n", err)
        return
    }
    fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 多行查询
// 多行查询db.Query()执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。
// 查询多条数据示例
func queryMultiRowDemo() {
    sqlStr := "select id, name, age from user where id > ?"
    rows, err := db.Query(sqlStr, 0)
    if err != nil {
        fmt.Printf("query failed, err:%v\n", err)
        return
    }
    // 非常重要：关闭rows释放持有的数据库链接
    defer rows.Close()

    // 循环读取结果集中的数据
    for rows.Next() {
        var u user
        err := rows.Scan(&u.id, &u.name, &u.age)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            return
        }
        fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
    }
}

// 插入数据
// 插入、更新和删除操作都使用方法。
// Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。
// 插入数据
func insertRowDemo() {
    sqlStr := "insert into user(name, age) values (?,?)"
    ret, err := db.Exec(sqlStr, "王五", 38)
    if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
    }
    theID, err := ret.LastInsertId() // 新插入数据的id
    if err != nil {
        fmt.Printf("get lastinsert ID failed, err:%v\n", err)
        return
    }
    fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
    sqlStr := "update user set age=? where id = ?"
    ret, err := db.Exec(sqlStr, 39, 3)
    if err != nil {
        fmt.Printf("update failed, err:%v\n", err)
        return
    }
    n, err := ret.RowsAffected() // 操作影响的行数
    if err != nil {
        fmt.Printf("get RowsAffected failed, err:%v\n", err)
        return
    }
    fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
    sqlStr := "delete from user where id = ?"
    ret, err := db.Exec(sqlStr, 3)
    if err != nil {
        fmt.Printf("delete failed, err:%v\n", err)
        return
    }
    n, err := ret.RowsAffected() // 操作影响的行数
    if err != nil {
        fmt.Printf("get RowsAffected failed, err:%v\n", err)
        return
    }
    fmt.Printf("delete success, affected rows:%d\n", n)
}

// 预处理查询示例
func prepareQueryDemo() {
    sqlStr := "select id, name, age from user where id > ?"
    stmt, err := db.Prepare(sqlStr)
    if err != nil {
        fmt.Printf("prepare failed, err:%v\n", err)
        return
    }
    defer stmt.Close()
    rows, err := stmt.Query(0)
    if err != nil {
        fmt.Printf("query failed, err:%v\n", err)
        return
    }
    defer rows.Close()
    // 循环读取结果集中的数据
    for rows.Next() {
        var u user
        err := rows.Scan(&u.id, &u.name, &u.age)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            return
        }
        fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
    }
}


// 预处理插入示例
func prepareInsertDemo() {
    sqlStr := "insert into user(name, age) values (?,?)"
    stmt, err := db.Prepare(sqlStr)
    if err != nil {
        fmt.Printf("prepare failed, err:%v\n", err)
        return
    }
    defer stmt.Close()
    _, err = stmt.Exec("小王子", 18)
    if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
    }
    _, err = stmt.Exec("沙河娜扎", 18)
    if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
    }
    fmt.Println("insert success.")
}

// 事务操作示例
func transactionDemo() {
    tx, err := db.Begin() // 开启事务
    if err != nil {
        if tx != nil {
            tx.Rollback() // 回滚
        }
        fmt.Printf("begin trans failed, err:%v\n", err)
        return
    }
    sqlStr1 := "Update user set age=30 where id=?"
    _, err = tx.Exec(sqlStr1, 2)
    if err != nil {
        tx.Rollback() // 回滚
        fmt.Printf("exec sql1 failed, err:%v\n", err)
        return
    }
    sqlStr2 := "Update user set age=40 where id=?"
    _, err = tx.Exec(sqlStr2, 4)
    if err != nil {
        tx.Rollback() // 回滚
        fmt.Printf("exec sql2 failed, err:%v\n", err)
        return
    }
    err = tx.Commit() // 提交事务
    if err != nil {
        tx.Rollback() // 回滚
        fmt.Printf("commit failed, err:%v\n", err)
        return
    }
    fmt.Println("exec trans success!")
}



func main() {
    err := initDB() // 调用输出化数据库的函数
    if err != nil {
        fmt.Printf("init db failed,err:%v\n", err)
        return
    }
}



/*

CREATE DATABASE sql_test;

use sql_test;

CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

*/