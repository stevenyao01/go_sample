package main

import (
	"fmt"
	"github.com/go_sample/src/hash/mHash"
	"os"
)

func main() {

	key := ""
	id := 0
	name := ""
	var hashTable hash.HashTable
	for {
		fmt.Println("==========员工菜单==========")
		fmt.Println("insert 表示添加员工")
		fmt.Println("show   表示显示员工")
		fmt.Println("find   表示查询员工")
		fmt.Println("exit   表示退出员工")
		fmt.Println("请输入你的选择：")
		fmt.Scanln(&key)
		switch key {
		case "insert":
			fmt.Println("请输入员工id：")
			fmt.Scanln(&id)
			fmt.Println("请输入员工名字：")
			fmt.Scanln(&name)
			emp := &hash.Emp{
				ID:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.Show()
		case "find":
			fmt.Println("请输入你要查找的id：")
			fmt.Scanln(&id)
			emp := hashTable.Find(id)
			if emp == nil {
				fmt.Printf("id=%d的员工不存在\n", id)
			} else {
				//显示雇员信息
				emp.ShowMe()

			}
		case "exit":
			os.Exit(0)
		}
	}
}
