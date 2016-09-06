package main
import (
    "fmt"
)
func main(){
    x,y := 1,2
    a := [3]int{1,2}
    b := [3]int{1,2,0}
    a[2] = 3
    fmt.Println(x == y)
    fmt.Println(a == b)
    p := new ([10]int)
    fmt.Println(p)
    fmt.Println(a)
    
    c := [3][3]int{
        {1,1,1},
        {2,2,2},
        // {1:2,3:2,0:2},
    }
    c[2][1] = 2
    fmt.Println(c)
    
    for i:=0;i<len(c);i++ {
        for j:=0;j<len(c[i]);j++ {
            fmt.Println(c[i][j])
        }
    }
    
    for i:=0; i<3;i++ {
        v := 1
        fmt.Println(&v)
    }
}