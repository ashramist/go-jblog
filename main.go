package main

import "fmt"

func getPersonInfo() (string, int) {
	return "张三", 10
}
func main() {
	// 常量
	const MAX_SCORE = 100

	count := 10
	
	is_ok := true
	name, age := getPersonInfo()

	// 数组类型
	if is_ok {
		fmt.Printf("name = %s,age = %d\n", name, age)
	}

	fmt.Println("hello world")

	fmt.Printf("name = %s,age = %d", name, age)

}
