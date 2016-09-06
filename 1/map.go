package main
import (
    "fmt"
    // "sort"
)
func main(){
    //map比使用线性搜索要快很多，但是比索引满慢100倍
    // var m map[int]string = make(map[int]string)
    // // m := make(map[int]string)//简洁方式
    // m[1] = "OK"
    // // delete(m,1)
    // a := m[1]
    // fmt.Println(a)
    
    //每一级的map都需要单独进行make初始化
    // m := make(map[int]map[int]string)
    // m[1] = make(map[int]string)//这里需要二级才可以
    // m[1][1] ="OK"
    // a:=m[1][1]
    // fmt.Println(a)
    
    // a, ok := m[2][1]
    // if !ok {
    //     m[2]=make(map[int]string)
    // }
    // m[2][1] = "GOOD"
    // a=m[2][1]
    // // a, ok := m[2][1]//替换上句可使ok显示true
    // fmt.Println(a,ok)
    
    //迭代操作
    //说明
    //比foreach更高级一点
    // for i,v:=range slice{//对slice迭代,i计数器,索引整型,v值的拷贝
    //     slice[i]//可对slice进行操作，也可操作v不影响slice的值
    // }
    // for k,v := range map{//对map操作，返回键值对
    //     map[k]//可操作map，也可以操作k,v不影响map本身
    // }
    
    
    // sm := make([]map[int]string,5)
    //这种操作的是拷贝
    // for _,v := range sm {
    //     v = make(map[int]string,1)
    //     v[1]="OK"
    //     fmt.Println(v)
    // }
    // fmt.Println(sm)
    //这种是操作值
    // for i := range sm {
    //     //这种是操作值
    //     sm[i] = make(map[int]string,1)
    //     sm[i][1] = "OK"
    //     fmt.Println(sm[i])
    // }
    // fmt.Println(sm)
    
    // m:=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
    // s:=make([]int,len(m))
    // i:=0//因为此处m没有索引，而是key，所以自己实现计数器
    // for k,_ := range m {
    //     s[i] =k
    //     i++
    // }
    // sort.Ints(s)
    // fmt.Println(s)
    
    m1:=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
    fmt.Println(m1)
    m2:=make(map[string]int,len(m1))
    for k,v := range m1 {
        m2[v]=k
    }
    fmt.Println(m2)
}