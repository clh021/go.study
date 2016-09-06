package main
import (
    "fmt"
)

func main(){
    
    //defer
    //即使发生严重错误也会执行
    //支持匿名函数调用，常用于资源清理，文件关闭，解锁以及记录时间
    // fmt.Println("a")
    // defer fmt.Println("b")
    // defer fmt.Println("c")
    //a c b
    
    // for i := 0; i < 3; i++ {
    //     defer fmt.Println(i)
    // }//2 1 0
    for i := 0; i < 3; i++ {
        defer func(){fmt.Println(i)}()//3 3 3
        //闭包，一直引用i，此处是地址引用，最后return的是最后结果
    }
    
    A()
    B()
    C()
    
    
    var fs = [4]func(){}
    
    for i := 0;i<4;i++{
        defer fmt.Println("defer i=",i)
        defer func(){fmt.Println("defer_closure i=",i)}()
        fs[i] = func(){fmt.Println("closure i=",i)}
    }
    for _,f := range fs{
        f()
    }
}

func A() {
    fmt.Println("Func A")
}
func B() {
    defer func() {//defer需要放在 panic之前才有效
        if err := recover(); err != nil {
            fmt.Println("Recover in B")
        }
    }()
    // fmt.Println("Func B")
    panic("Panic in B")
    
}
func C() {
    fmt.Println("Func C")
}