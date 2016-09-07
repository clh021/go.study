package main
import (
    "io"
    "log"
    "net/http"
)
func main() {
    //设置路由
    // http.HandleFunc("/",handler)//访问具体的路径，使用哪个handler
    http.HandleFunc("/",sayHello)//降将sayHello
    err := http.ListenAndServe(":8080",nil)//两个参数:addr,handler
    if err!=nil {
        log.Fatal(err)//打印出错误
    }
}
func sayHello(w http.ResponseWriter, r *http.Request){//作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
    io.WriteString(w, "Hello world,this is version2")//输出的字符串
}