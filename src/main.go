package main

func main() {

}

func outer() func() int {
	var i = 10

	// 逃逸 i 变量, 使其能被外部访问
	return func() int {
		return i + 1
	}
}
