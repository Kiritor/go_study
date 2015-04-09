package main

import (
    "fmt"
    "github.com/ziutek/mymysql/mysql"
    //_ "github.com/ziutek/mymysql/native" // 线程不安全
    _ "github.com/ziutek/mymysql/thrsafe" // 线程安全
)

func main() {
    db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "root", "webdemo")

    err := db.Connect()
    if err != nil {
        panic(err)
    }

    rows, result, err := db.Query("select * from user")
    if err != nil {
        panic(err)
    }

    for _, row := range rows {
        fmt.Println(row.Str(0)+":"+row.Str(1)+":"+row.Str(2))
    }
	fmt.Println(result.AffectedRows())
	fmt.Println(result.Message())
	ins,err:=db.Prepare("insert into user values(?,?,?)")
	if err!=nil {
		fmt.Println(err)
	}
	//开启事务,db处于lock状态,只有当事务提交或者回滚之后才会解锁
	tr,_:=db.Begin()
	//处于事务且线程安全
	/**
      start方法属于db,因此下面两条插入语句都不会执行
    */
	go func(){
		tr.Start("insert user valies(50,'LCore','LCore')")
	}()
	tr.Start("insert user valies(70,'LCore','LCore')")
	tr.Do(ins).Run(60,"LCore","LCore")
	tr.Commit()   //事务提交
}
