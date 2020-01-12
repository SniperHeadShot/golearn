package base

import (
	"fmt"
	"math/rand"
	"time"
)

type user struct {
	Name string
	age  int
}

// 获取不同类型变量的默认值
func TypeDefaultValue() {
	var p1 bool
	var p2 int
	var p3 float32
	var p4 string
	var p5 user
	var p6 *int
	var p7 []int
	var p8 map[string]int
	var p9 chan int
	var p10 func(string) int
	var p11 error // error 是接口

	fmt.Printf("bool = %t,&bool = %v \n", p1, &p1)
	fmt.Printf("int = %d,&int = %v \n", p2, &p2)
	fmt.Printf("float32 = %f,&float32 = %v \n", p3, &p3)
	fmt.Printf("string = %s,&string = %v \n", p4, &p4)
	fmt.Printf("struct = %+v,&struct = %p \n", p5, &p5)
	fmt.Printf("*int = %v,&*int = %p \n", p6, &p6)
	fmt.Printf("[]int = %v,&[]int = %p \n", p7, &p7)
	fmt.Printf("map[string] int = %v,&map[string] int = %p \n", p8, &p8)
	fmt.Printf("chan int = %v,&chan int = %p \n", p9, &p9)
	fmt.Printf("&func(string) int = %p \n", &p10)
	fmt.Printf("interface = %v,&interface = %p \n", p11, &p11)
}

// const iota 关键字的使用
func ConstIota() {
	const (
		a = iota      //0
		b             //1
		c             //2
		d = "ha"      //独立值，iota += 1
		e             //"ha"   iota += 1
		f = 100       //iota +=1
		g             //100  iota +=1
		h = iota      //7,恢复计数
		i = 1 << iota //1<<8
		j             //1<<9
		k             //1<<10
		l             //1<<11
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
	// 结论：iota 会不断自增，如果遇到iota参与了计算，则下一项依然会使用之前的表达式运算
}

// switch fallthrough break 关键字使用
func SwitchFallthrough() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := r.Intn(6)
	switch randomInt {
	// 多条件匹配
	case 1, 2:
		fmt.Println("case 1, 2")
		// 不同的 case 之间不使用 break 分隔，默认只会执行一个 case
		break
	case 3:
		fmt.Println("case 3")
	case 4:
		fmt.Println("case 4")
		// 用 fallthrough 关键字可以执行多个 case
		fallthrough
	case 5:
		fmt.Println("case 5")
	default:
		fmt.Printf("default case %d \n", randomInt)
	}
}

// 遍历
func Traversal() {
	arr := []string{"abc", "def", "ghi"}

	// 方式一
	index := 0
	for index < len(arr) {
		if index == 2 {
			break
		}
		fmt.Printf("method1 ==> index = %d, elemnet = %s \n", index, arr[index])
		index++
	}

	// 方式二
	for i := 0; i < len(arr); i++ {
		if i == 1 {
			continue
		}
		fmt.Printf("method2 ==> index = %d, elemnet = %s \n", i, arr[i])
	}
loop:
	// 方式三
	for index, element := range arr {
		if index == 3 {
			goto loop
		}
		fmt.Printf("method3 ==> index = %d, elemnet = %s \n", index, element)
	}
}

// 数组初始化
func InitArr() {
	var arr1 []string
	var arr2 [3]string
	fmt.Printf("arr1=%v, len(arr1)=%d,arr2=%v, len(arr2)=%d", arr1, len(arr1), arr2, len(arr2))
}
