package main

import "fmt"

func main() { // { 不能在单独的行上
	//{  // 错误，{ 不能在单独的行上
	fmt.Printf("hello world \n")

	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}
