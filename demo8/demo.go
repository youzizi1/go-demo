package main
import "fmt"
func main()  {
	// 创建空切片
	s := make([]string, 3)
	fmt.Println(s)
	// 赋值
	s[1] = "hello"
	fmt.Println(s, s[1], len(s))
	// 追加
	s = append(s, "world")
	fmt.Println(s, s[3], len(s))
	// 复制
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	// 剪切
	m := make([]int, 4)	
	m[0] = 1
	m[1] = 2
	m[2] = 3
	m[3] = 4
	n := m[1:3]
	fmt.Println(n)
	n = m[:3]
	fmt.Println(n)
	n = m[1:]
	fmt.Println(n)

	t := []string{"a", "b", "c"}
	fmt.Println(t, t[1:])
}