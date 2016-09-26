package main

import (
	"io"
    "io/ioutil"
	"log"
	"net/http"
	"reflect"//用来获取对象的类型
	"encoding/json"//
	"os"
	"fmt"//用来终端调试
)

/*
 * 这里是尝试写一个监控程序，有如下效果：
 * 1、可以实时看到举报同步的情况
 * 2、可以监测到服务器部分各个服务是否运转正常
 * 3、可以监测到4G系统部分各个功能模块是否运转正常
 * -----------------------------------------------
 * 您可以通过该命令运行起来并记录日志: nohup server >> /tmp/nohup.out 2>&1 & 
 */

func main() {
    // http.HandleFunc("/",Homepage)//降将sayHello
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
    http.HandleFunc("/heartbeat",Heartbeat)//降将sayHello
 	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd+""))))
    err = http.ListenAndServe(":8081",nil)//两个参数:addr,handler
    if err!=nil {
        log.Fatal(err)//打印出错误
    }
    
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/heartbeat", Heartbeat)
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//  	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd+""))))
//     mux.HandleFunc("/",Homepage)    
// 	err = http.ListenAndServe(":8081", mux) //两个参数:addr,handler
// 	if err != nil {
// 		log.Fatal(err) //打印出错误
// 	}
}
type HeartbeatItem struct {
    Server 	string
    System 	string
    Jbsync 	string
}

func Show(item HeartbeatItem, w http.ResponseWriter) {
	fmt.Println(reflect.TypeOf(item.Server))
	io.WriteString(w,item.Server+"<hr />")
	io.WriteString(w,item.System+"<hr />")
	io.WriteString(w,item.Jbsync+"<hr />")
}
func Heartbeat(w http.ResponseWriter, r *http.Request)() { //作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
    s := HeartbeatItem{GetServer(),GetSystem(),GetJbsync()}
    res, err := json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    } else {
    	w.Header().Set("Content-Type", "application/json")
    }
    // fmt.Println(reflect.TypeOf(res))
    // fmt.Println(string(res))
    
	io.WriteString(w, string(res)) //输出的字符串
}
//获取服务器各个状态信息
func GetServer()(resp string) {
    //包含诸如：内存，网络，磁盘，CPU，redis,ftp,mysql,ping等服务的状态
    resp="服务器的各个系统状态将会呈现在这里<br />"
    return
}
//尝试访问4G系统部分各个功能模块以确定运作是否正常
func GetSystem()(resp string) {
    //包含诸如：api,app,4g,jubaosync,
    resp="4G系统各个功能模块的运作结果将会呈现在这里<br />"
    return
}
//得到举报功能状态，以及举报详情，同步详情
func GetJbsync()(resp string) {
    //包含诸如：举报时间，同步时间，举报ID，上次同步状态，本次同步状态
	url:="http://4g.womenxing.com/?r=site/syncjubao3"
	resp, statusCode := Get(url)
	if 200 != statusCode {
		resp = ""
	}
	return
}

func Get(url string) (content string, statusCode int) {
    resp, err1 := http.Get(url)
    if err1 != nil {
        statusCode = -100
        return
    }
    defer resp.Body.Close()
    data, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        statusCode = -200
        return
    }
    statusCode = resp.StatusCode
    content = string(data)
	// fmt.Println(reflect.TypeOf(content))
	// fmt.Println(content)
	// fmt.Println(statusCode)
    return
}
func Homepage(w http.ResponseWriter, r *http.Request) () {
    content:=""
    io.WriteString(w,content)
}