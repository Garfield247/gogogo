package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello world !"
	// 是否包含
	fmt.Println(strings.Contains(s1, "he")) // true
	// 是否包含任意一个
	fmt.Println(strings.ContainsAny(s1, "abcde")) // true
	// 出现的次数
	fmt.Println(strings.Count(s1, "l")) // 3
	// 是否以什么开头
	fmt.Println(strings.HasPrefix(s1, "h")) // true
	// 是否以什么结尾
	fmt.Println(strings.HasSuffix(s1, "!")) // true
	//返回索引 不存在返回 -1
	fmt.Println(strings.Index(s1, "l")) // 2
	// 任意一个字符出现的位置
	fmt.Println(strings.IndexAny(s1, "abcdeh")) // 0
	// 最后一次出现的位置
	fmt.Println(strings.LastIndex(s1, "l")) // 9
	// 拼接String
	ss1 := []string{"adsad", "dasdasca", "dasdas"}
	fmt.Println(strings.Join(ss1, "--"))

	// 分割字符串
	s2 := "adsa0saddsa-dasd-dsadasd-asdasd-asd-asdds"
	ss2 := strings.Split(s2, "-")
	fmt.Printf("%T-%v\n", ss2, ss2)

	// 重复拼接自身
	fmt.Println(strings.Repeat("abs-", 3))

	// 字符串的替换 ,传入-1 全部替换
	s3 := strings.Replace(s2, "a", "A", -1)
	fmt.Println(s3)
	// 转换大小写
	fmt.Println(strings.ToLower(s3))
	fmt.Println(strings.ToUpper(s3))
}
