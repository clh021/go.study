package main
import (
    "io"
    "log"
    "net/http"
    "os"
)
func main() {
    mux := http.NewServeMux()//上例中http.HandleFunc其实就是一个ServeMux//更加底层的操作//自己实现
    // mux.Handle(pattern, handler)
    mux.Handle("/",&myHandler{})
    mux.HandleFunc("/hello",sayHello)
    
    //一个静态文件服务部分
    wd,err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    mux.Handle("/static/",
        http.StripPrefix("/static/",http.FileServer(http.Dir(wd))))
    //尝试访问一下以下地址
    //localhost:8080/static/static/1.txt
    //localhost:8080/static/static/
    //localhost:8080/static/
    
    //err 前面有 则去掉 := 中的 :
    err = http.ListenAndServe(":8080",mux)//两个参数:addr,handler
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