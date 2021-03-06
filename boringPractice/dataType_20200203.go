package main

import "fmt"

// struct
type animal struct {
	dna string
	age int
}
type dog struct {
	animal // 內嵌 animal
	hiar   string
}
type cat struct {
	animal animal // 結構再包結構animal
	behave string
}

// method :定義屬於type的func
func (d *dog) growth(year int) {
	d.age += year
}
type myInt int // 不是結構也可以
func (i *myInt) add(num myInt) {
	*i += num
}

func main() {
	// struct 範例
	myDog := dog{
		animal: animal{"dogDna", 2},
		hiar:   "dog hair",
	}
	fmt.Println(myDog, myDog.age, myDog.animal)

	myCat := cat{
		animal: animal{"catDna", 1},
		behave: "mew~",
	}
	fmt.Println(myCat, myCat.animal.age, myCat.animal)

	// func :在go種function也是一種型別可以assign給變數
	func1 := func(msg string) { fmt.Println("func1", msg) }
	func1("hi go func")

	// array :陣列, 大小固定
	strs := [2]string{"hello", "world"}
	ns := [...]int{1, 2, 3, 5, 6, 8, 9, 10}
	fmt.Println(strs, ns, len(ns))
	for i, j := range strs {
		fmt.Println(i, j)
	}

	// slice :可動態改變大小, slice傳至function本身是by refer
	s1 := ns[2:5] // 從別人那邊切來
	fmt.Println(s1, len(s1), cap(s1))
	s2 := []string{"hello", "world"} // 直接宣告
	fmt.Println(s2, len(s2), cap(s2))

	s2 = append(s2, "new", "slice")
	fmt.Println(s2, len(s2), cap(s2))

	// interface{} :空介面可以assign不同的資料型態
	var x interface{}
	nums := []int{1, 2, 3}
	x = nums
	fmt.Println(x)

	// map :可以定義key-value
	m1 := map[string]int{
		"str1": 1,
		"str2": 3,
	}
	fmt.Println(m1)

	age, ok := m1["age"] // 檢查map的key是否存在
	if ok == true {
		fmt.Println(age)
	} else {
		fmt.Println(`not found "age"`)
	}
	m1["newk"] = 7 // 新增key-value
	fmt.Println(m1)

	delete(m1, "str2") // 刪除特定key
	fmt.Println(m1)

	fmt.Println()

	// 基本型態與定義
	//basicType()
}

func basicType() {
	// 基本型別: int, float, boolean, string; 也可用const宣告常數
	var num int
	var fnum float32
	var flag bool
	var str string
	fmt.Println(num, fnum, flag, str)

	// 指標: 宣告 *type ; 取值 *var ; 取址 &var
	var pt1 *int
	pt1 = &num
	*pt1 = 5
	fmt.Println(pt1, *pt1, &num, num)

	// 別名: 為型態取別名, 內建的別名有: type byte = uint8 ; type rune = int32
	type myInt = int
	var mynum myInt
	mynum = 7
	fmt.Printf("%d %T\n", mynum, mynum) // %T格式化輸出: 查看型態

	// 新型別: 自定義新型別
	type myFloat float32
	var myflt myFloat
	myflt = 7.2
	fmt.Printf("%f %T\n", myflt, myflt)
}
