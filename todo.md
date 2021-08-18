1. go path已经弃用，建议使用go module
2. go会自动推断已经初始化的变量类型
3. 若声明变量但未赋值，会初始化为零值。例如：int的零值为0
4. 函数内可以通过break或者return跳出循环
5. go里面没有三目运算符，只能使用if
6. 没有表达式的switch是实现ifelse逻辑的另外一种方式
7. fmt.Println打印数组的时候会按照[v1, v2, ...]显示
8. import 一个未使用的包或者一个定义但未使用的变量，编译会报错
9. cannot find module for path xxx报错可以通过go env -w GO111MODULE=auto来解决
10. const声明的变量必须首字母大写