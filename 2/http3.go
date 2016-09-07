package main
import (
    "io"
    "log"
    "net/http"
)
func main() {
    mux := http.NewServeMux()//上例中http.HandleFunc其实就是一个ServeMux//更加底层的操作//自己实现
    // mux.Handle(pattern, handler)
    mux.Handle("/",&myHandler{})
    mux.HandleFunc("/hello",sayHello)
    err := http.ListenAndServe(":8080",mux)//两个参数:addr,handler
    if err!=nil {
        log.Fatal(err)//打印出错误
    }
}
type myHandler struct{}
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){//作为要注册路由的hanlder
    //具体有哪些参数和参数类型等要求看文档
    io.WriteString(w, "URL:"+r.URL.String())//输出的字符串
}
func sayHello(w http.ResponseWriter, r *http.Request){//作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
    io.WriteString(w, "Hello world,this is version3")//输出的字符串
}