package main

import (
	"fmt"
)

/**
 * @Package Name: surgemq
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-6 上午11:42
 * @Description:
 */


type LList struct {
	val  int
	next  *LList
}

func printList(str *LList) {
	for str != nil {
		fmt.Println(*str)
		str = str.next
	}
}

func main() {
	var lista LList
	var listb LList
	initList(&lista, 2, 4, 3)

	initList(&listb, 4, 6, 5)

	listn := addTwoNumbers(lista, listb)

	printList(&listn)
}

func addTwoNumbers(lista LList, listb LList)(LList){
	aFirValue, aMidValue, aTailValue := getValue(lista)

	bFirValue, bMidValue, bTailValue := getValue(listb)

	// 取出拼为第一个三位数
	aValue := aFirValue*100 + aMidValue*10 + aTailValue
	// 拼第二个三位数
	bValue := bFirValue*100 + bMidValue*10 + bTailValue

	// 总和
	total := aValue + bValue

	nT := total - total%100
	//百位
	newFirValue := nT/100
	//十位
	newMidValue := (total - nT)/10
	//个位
	newTaiValue := (total -nT)%10

	var listn LList
	initList(&listn, newFirValue, newMidValue, newTaiValue)
	return listn
}


func initList(listNew1 *LList, a int, b int, c int)() {
	//var listNew1 LList
	var listNew2 LList
	var listNew3 LList

	listNew1.val = a
	listNew1.next = &listNew2

	listNew2.val = b
	listNew2.next = &listNew3

	listNew3.val = c
}

func getValue(list LList)(int, int, int){
	return list.val, list.next.val, list.next.next.val
}