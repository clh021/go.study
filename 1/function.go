
    //Go function 不支持嵌套,重载,默认参数
    //支持 无需原型,不定长度变参,多返回值,命名返回值,匿名函数,闭包
    // H("OK",1,2)
    a, b := 1, 2
    I(a,b)
    fmt.Println(a, b)//即使I函数改变值，也不会影响
    
    s1 := []int{1,2}
    J(s1)//
    fmt.Println(s1)//直接传递slice会传递内存地址，函数修改值会有影响
    
    s := 9
    K(&s)//指针传递
    fmt.Println(s)
    
    l := L//L is a func//这里是func类型//Go语言一切皆类型
    l()
    // fmt.Println(l)
    
    //匿名函数
    
    m := func() {
        fmt.Println("Func no name")
    }
    m()
    
    //闭包函数调用
    // clo := closure(12)
    // fmt.Println(clo(1))
    // fmt.Println(clo(4))
    
    
}

//函数
// func A(){
// }

// func B(a int,b string,c int){
// }

// func C(a,b,c int){
// }

// func D()(a,b,c int){//返回值
// }

// func E()(a,b,c int){//返回值
// }

// func F()(a,b,c int){//返回值
//     a,b,c = 1,2,3
//     return//会自动返回a,b,c
// }

// func G()(int,int,int){//返回值需要指定
//     a,b,c = 1,2,3
//     return a,b,c//强烈建议
// }


// func H(b string,a ...int) {//返回值
//     //不定长变参
//     //a 接收后为slice
//     fmt.Println(a)
// }

func I(s ...int) {
    s[0]=3
    s[1]=4
    fmt.Println(s)//是值拷贝，不会影响外部值
}

func J(s []int) {
    s[0]=3
    s[1]=4
    fmt.Println(s)//是slice地址拷贝，会影响外部值
}

func K(a *int) {//引用传递//指针类型
    *a=4//去除取出值来进行操作
    fmt.Println(*a)//是slice地址拷贝，会影响外部值
}

func L() {
    fmt.Println("Func L")
}

func closure(x int) func(int) int {//闭包函数返回一个函数
    //计划返回一个匿名函数
    fmt.Printf("%p\n",&x)//打印x内存地址
    return func(y int) int{
        fmt.Printf("%p\n",&x)//打印x内存地址//x是同一个值同一个地址
        return x + y
    }
}