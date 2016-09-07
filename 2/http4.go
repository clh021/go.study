package main
import (
    "io"
    "log"
    "net/http"
    "time"
)
var mux map[string]func(http.ResponseWriter, *http.Request)
func main() {
    server := http.Server{
        Addr: ":8080",
        Handler: &myHandler{},
        ReadTimeout: 5*time.Second,
    }
    
    //如何绑定自己的sayHello呢？
    mux = make(map[string]func(http.ResponseWriter, *http.Request))//上面要 var mux
    //此处可直接复制，相当于申请一点内存空间
    mux["/hello"] = sayHello
    mux["/bye"] = sayBye
    //serveHttp中进行map
    //---end
    
    err := server.ListenAndServe()//两个参数:addr,handler
    if err!=nil {
        log.Fatal(err)//打印出错误
    }
}
type myHandler struct{}
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){//作为要注册路由的hanlder
    if h,ok := mux[r.URL.String()]; ok {
        h(w,r)
        return
    }
    //具体有哪些参数和参数类型等要求看文档
    io.WriteString(w, "URL:"+r.URL.String())//输出的字符串
}

func sayHello(w http.ResponseWriter, r *http.Request){//作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
    io.WriteString(w, "Hello world,this is version3")//输出的字符串
}

func sayBye(w http.ResponseWriter, r *http.Request){//作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
    io.WriteString(w, "Bye world,this is version3")//输出的字符串
}