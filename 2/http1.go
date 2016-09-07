package main
import (
    "github.com/astaxie/beego"
)
type HomeController struct { // 定义自己的结构
    beego.Controller //嵌入beego自带的方法，就可直接使用了
}
//重载 get方法，即可隐藏beego的get方法，使用我们的get方法
func (this *HomeController) Get() {
    this.Ctx.WriteString("Hello World")
}
func main() {
    beego.Router("/",&HomeController{})//路径,要注册的Controller
    beego.Run()//启动服务
}