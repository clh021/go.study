package main
import (
    "fmt"
)
func main(){
    // var s1 []int //这是一个空的slice
    // a := [10]int{1,2,3,4,5,6,7,8,9,0}
    // s1 := a[9] //取数组中索引为9的元素
    // s1 := a[5:10] //截取数组中后面的5个元素
    // s1 := a[5:len(a)] //截取数组中第5个后面的所有元素
    // s1 := a[5:] //截取数组中第5个后面的所有元素
    // s1 := a[:5] //截取数组中去前面的5个元素
    // s1 := a[0:5] //截取数组中去前面的5个元素
    // fmt.Println(s1)
    
    // s1:=make([]int,3,10)//包含3个int元素,10小块连续的内存,之后不用重新分配内存会更快
    //能够实现事先知道内存需要多少，可以避免多次重新分配内存，效率会高很多
    // fmt.Println(s1)
    
    // a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
    // sa:=a[2:5]
    // fmt.Println(sa)//[99 100 101]
    // fmt.Println(string(sa))//cde
    // fmt.Println(len(a),len(sa),cap(a),cap(sa))//11(元素个数) 3(元素个数)  11(容量) 9(容量)
    
    //reslice
    // sb:=sa[1:3]
    // fmt.Println(sb)
    // fmt.Println(string(sb))//de
    
    //append
    // s1 := make([]int, 3, 6)
    // fmt.Printf("%p\n",s1)//第一个内存地址
    // s1=append(s1,1,2,3)
    // fmt.Printf("%v %p\n",s1,s1)//第二个内存地址
    
    // a := []int{1,2,3,4,5}
    // s1 := a[2:5]
    // s2 := a[1:3]
    // fmt.Println(s1,s2)//
    // // s2 = append(s2,1,1,1,1,1,1,1,1,1,1)//若超过了容量则重新分配即不会影响
    // s1[0]=9//若同内存地址，则一个变，都变，若不同，则不变
    // fmt.Println(s1,s2)//
    
    //copy //拷贝结果个数以被拷贝对象为准
    s1 := []int{1,2,3,4,5,6}
    s2 := []int{7,8,9}
    copy(s1,s2)
    fmt.Println(s1)//[7 8 9 4 5 6]
    
    s3 := []int{1,2,3,4,5,6}
    s4 := []int{7,8,9}
    copy(s4,s3)
    fmt.Println(s4)//[1 2 3]
    
    s5 := []int{1,2,3,4,5,6}
    s6 := []int{7,8,9,10,1,1,1,1,1}
    copy(s6[2:4],s5[1:3])
    fmt.Println(s6)//[7 8 2 3 1 1 1 1 1]
}