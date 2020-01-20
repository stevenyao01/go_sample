package main;

import (
	"fmt"
	"database/sql"
	_ "code.google.com/p/odbc"
)

func main() {
	conn, err := sql.Open("odbc", "driver={Microsoft Access Driver (*.mdb,*.accdb)};dbq=/home/steven/201904.mdb");
	if (err != nil) {
		fmt.Println("Connecting Error");
		return;
	}
	defer conn.Close();
	stmt, err := conn.Prepare("select top 5 id from ab_contents");
	if (err != nil) {
		fmt.Println("Query Error", err);
		return;
	}
	defer stmt.Close();
	row, err := stmt.Query();
	if err != nil {
		fmt.Println("Query Error", err);
		return;
	}
	defer row.Close();
	for row.Next() {
		var id int;
		if err := row.Scan(&id); err == nil {
			fmt.Println(id);
		}
	}
	fmt.Printf("%s\n", "finish");
	return;
}


//package main;
//import (
//	"fmt"
//	"database/sql"
//	//_ "code.google.com/p/odbc"
//	_ "github.com/weigj/odbc"
//)
//
//
//
////package main;
////import (
////	"fmt"
////	"database/sql"
////	_"odbc/driver"
////	//_ "github.com/weigj/odbc"
////)
//
///**
// * @Package Name: odbc
// * @Author: steven yao
// * @Email:  yhp.linux@gmail.com
// * @Create Date: 19-4-27 下午7:21
// * @Description:
// */
//
//
//
////
////func main(){
////	conn,err := sql.Open("odbc","driver={Microsoft Access Driver (*.mdb,*.accdb)};dbq=/home/steven/test.mdb");
////	if(err!=nil){
////		fmt.Println("Connecting Error ", err);
////		return;
////	}
////	defer conn.Close();
////	stmt,err := conn.Prepare("select * from EXTRAINFO");
////	if(err!=nil){
////		fmt.Println("prepare Error ", err);
////		return;
////	}
////	defer stmt.Close();
////	row,err := stmt.Query();
////	if err!=nil {
////		fmt.Println("Query Error");
////		return;
////	}
////	defer row.Close();
////	for row.Next() {
////		var id int;
////		var name string;
////		if err := row.Scan(&id,&name);err==nil {
////			fmt.Println(id,name);
////		}
////	}
////	fmt.Printf("%s\n","finish");
////	return;
////}
//
//
//func main() {
//	//conn := fmt.Sprintf("driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=D:/TIG/Scale_Run.MDB;Uid=;Pwd=***;")
//	conn := fmt.Sprintf("driver={Microsoft Access Driver (*.mdb,*.accdb)};server=127.0.0.1;dbq=/home/steven/201904.mdb")
//	db, err := sql.Open("odbc", conn)
//	if err != nil {
//		fmt.Println("数据库打开错误:", err)
//	}
//	defer db.Close()
//	var result string
//	//err = db.QueryRow("SELECT TOP 1 序号 FROM Run0 ORDER BY 序号 DESC").Scan(&result) //读取最新的标记
//	//err = db.QueryRow("SELECT * from EXTRAINFO BY IDNO DESC").Scan(&result) //读取最新的标记
//	err = db.QueryRow("SELECT * from EXTRAINFO BY IDNO DESC").Scan(&result) //读取最新的标记
//	if err != nil {
//		fmt.Println("数据库查询错误:", err)
//	}
//	fmt.Println(result)
//}
//
////
////func main(){
////	conn,err := sql.Open("odbc","driver={Microsoft Access Driver (*.mdb)};dbq=/home/steven/code/go/src/github.com/go_sample/201904.mdb");
////	if(err!=nil){
////		fmt.Println("Connecting Error ", err);
////		return;
////	}
////	defer conn.Close();
////	stmt,err := conn.Prepare("SELECT * from EXTRAINFO BY IDNO DESC");
////	if(err!=nil){
////		fmt.Println("Query Error ", err);
////		return;
////	}
////	defer stmt.Close();
////	row,err := stmt.Query();
////	if err!=nil {
////		fmt.Println("Query Error");
////		return;
////	}
////	defer row.Close();
////	for row.Next() {
////		var id int;
////		var name string;
////		if err := row.Scan(&id,&name);err==nil {
////			fmt.Println(id,name);
////		}
////	}
////	fmt.Printf("%s\n","finish");
////	return;
////}
//
////func main() {
////	//conn := fmt.Sprintf("driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=D:/TIG/Scale_Run.MDB;Uid=;Pwd=***;")
////	conn := fmt.Sprintf("driver={Microsoft Access Driver (*.mdb)};dbq=/home/steven/201904.mdb")
////	db, err := sql.Open("odbc", conn)
////	if err != nil {
////		fmt.Println("数据库打开错误:", err)
////	}
////	defer db.Close()
////	var result string
////	//err = db.QueryRow("SELECT TOP 1 序号 FROM Run0 ORDER BY 序号 DESC").Scan(&result) //读取最新的标记
////	//err = db.QueryRow("SELECT * from EXTRAINFO BY IDNO DESC").Scan(&result) //读取最新的标记
////	err = db.QueryRow("SELECT * from EXTRAINFO BY IDNO DESC").Scan(&result) //读取最新的标记
////	if err != nil {
////		fmt.Println("数据库查询错误:", err)
////	}
////	fmt.Println(result)
////}