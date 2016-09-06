package main
import (
    "fmt"
)
const x int =1
const y = 'A'
const str1,str2,str3 = "Hello, 世界","Hello, 陈良红","我后面有换行吗\n？"
const (
    i,m,j = iota,x,y
    o,p,q = iota,len(str2),len(str3)
)
const (
    r = iota
    s
    t
)
type (
    byte int8
)
func main(){
    // a,_,c,d := 1,2,3,4
    // fmt.Println(a)
    // fmt.Println(c)
    // fmt.Println(d)
// 	fmt.Println(str1)
// 	fmt.Println(str2)
// 	fmt.Println(str3)
// 	fmt.Println(len(str2))
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(i)
// 	fmt.Println(m)
// 	fmt.Println(j)
// 	fmt.Println(o)
// 	fmt.Println(p)
// 	fmt.Println(q)
// 	fmt.Println(r)
// 	fmt.Println(s)
// 	fmt.Println(t)
	
	LABEL1:
	for {
	    for i := 0; i < 10; i++{
	        if i > 3 {
	            // goto LABEL1//如果LABEL1放在循环的后面可以用goto
	            // continue LABEL1//会继续当前循环
	            fmt.Println(i)
	            break LABEL1
	        }
	    }
	}
	fmt.Println("OK")
}