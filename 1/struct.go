package main
import (
    "fmt"
)
type person struct {
    Name string
    Age int
    Contact struct {
        Phone, City string//类型放在变量名后面会很便利而且不会产生歧义
    }
}
type teacher struct {
    person
    Teacherid int
}
type student struct {
    person
    Studentid int
}
func main(){
    teacher := teacher{Teacherid:10001,person:person{Name:"Liu Laoshi",Age:38}}
    student := student{Studentid:10002,person:person{Name:"Liu shengqi",Age:28}}
    fmt.Println(teacher,student)
    teacher.Teacherid=10005
    teacher.Name="Niu Laoshi"
    teacher.person.Age=37
    // student.person.C//尝试设置学生的联系方式
    fmt.Println(teacher,student)
    
    a := person{
        Name:"leehom",
    }
    a.Age = 19
    fmt.Println(a)
    A(a)
    fmt.Println(a)
    B(&a)
    a.Contact.Phone = "121234567"
    a.Contact.City = "WuHan"
    fmt.Println(a)
    // C(&a)//如果有很多函数都需要指针传递可以直接在 a:=person的时候就可以取出指针地址 a:=person // 推荐对初始化结构都进行&
    // fmt.Println(a)
    
    
    b := &struct { //匿名结构
        Name string
        Age int
    }{
        Name:"leehom",
        Age:19,
    }
    fmt.Println(b)
    
    c := &struct {
        string
        int
    }{
        "lianghong"//匿名函数没有名字和字段名字，只能依赖顺序来进行赋值
        28
    }
    
    
    
}

func A (per person) {
    per.Age = 13
    fmt.Println("A",per)
}
func B (per *person) {
    per.Age = 13
    fmt.Println("B",per)
}
func C (per *person) {
    per.Age = 13
    fmt.Println("C",per)
}