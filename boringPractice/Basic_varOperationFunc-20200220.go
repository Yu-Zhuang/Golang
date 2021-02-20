package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello Go")
	fmt.Println(varGlobal1)
	
	fmt.Println(myFunc(1,"hi"))
	fmt.Println(myFunc(-3,"hola"))

	fmt.Println("end")
}

// how to announce variable in Go
var (
	varGlobal1 int = 5
	varGlobal2 = 3
	// 供外部檔可用的public 變數or函數(func)的第一個字要大寫
	RealGlobal = 7
)
// 宣告常數
const (
	vConst1 = 1
	vConst2 = 2
)

// variable announce
func varAnnounce(){
	// type 1 (this is the only method that can announce outside of func)
	var varName1, varName2 int = 1, 2
	// type 2
	var varName3 = 3
	// type 3
	varName4 := 4

	/*
	basic type of Go:
	1. int : 1
	2. float : 1.2
	3. string : "im string"
	4. bool : true / false
	5. null : nil
	ps. const : const
		匿名佔位符(anonymous placeholder): _
	*/
	// because var in func have to use, or you can't compile so we have this useless line
	varName1 = varName2 + varName3 + varName4 + varName1

	// 多重賦值應用 swap
	varName1, varName2 = varName2, varName1
}

// operation
func operation() {
	// loop
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
	// for as while loop
	i := 0
	for i<10 {
		fmt.Println(i)
		i ++
	}

	// if / else
	for {
		i ++
		fmt.Println(i)
		if i>20 {
			break
		} else {
			fmt.Println("not complete")
		}
	}
	// if operation ; judge { } 寫法
	for {
		fmt.Println(i)
		if i++; i>30 {
			break
		} else {
			fmt.Println("not complete2")
		}
	}	

	// switch
	switch i {
	case 29:
		fmt.Println("i == 29")
	case 30:
		fmt.Println("i == 30")
	case 31:
		fmt.Println("i == 31")
	default:
		fmt.Println("i != 29 && i != 30")
	}
}

// array
func array() {
	// int array
	var intAry1 [3]int
	var intAry2 = [3]int {1,2,3}
	// 有指定初始值可用[...], 用len()可取得array長度
	intAry3 := [...]int {4,5,6}

	fmt.Println(intAry1[2])
	fmt.Println(intAry2[2])
	fmt.Println(intAry3[0], len(intAry3))
	
	// string array
	var str1 [2]string
	str2 := [2]string { "hello", "Golang" }

	fmt.Println(str1)
	fmt.Println(str2)

	// travel in array
	for i, j := range str2 {
		fmt.Println(i, j)
	}

}

// 函數func funcName(varName varType) (retureVar retureType) {}
// func 也可被作為一種資料型別進行賦值與帶入函式
func myFunc(num int, str string) (int, bool) {
	value := 1

	if num > 0 {
		return value, true
	}
	return  value*-1, false
}