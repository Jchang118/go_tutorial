package main

import (
    myhttp "go_tutorial/go_basic/http"
    //"encoding/json"
    "fmt"
    //"html/template"
    "io"
    "net/http"
    "os"
    //"strconv"
    "strings"
    //"time"
)

// 详细看一下http协议
func HttpObservation(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("request method: %s\n", r.Method)
    fmt.Printf("request host: %s\n", r.Host) //服务端host
    fmt.Printf("request url: %s\n", r.URL)
    fmt.Printf("request proto: %s\n", r.Proto)
    fmt.Println("request header")
    for key, values := range r.Header {
        fmt.Printf("%s: %v\n", key, values)
    }
    fmt.Println()
    fmt.Printf("request body: ")
    io.Copy(os.Stdout, r.Body) //把r.Body流里的内容拷贝到os.Stdout流里
    fmt.Println()

    // 必须先设置响应头,再设置响应码,最后设置响应体,否则无效
    w.Header().Add("tRAce-id", "4723956498105") //在WriteHeader之前设置Header.header里的key是大小写不敏感的,会自动把每个单词(各单词用-连接)的首字母转为大写,其他字母转为小写
    w.WriteHeader(http.StatusBadRequest)    //设置StatusCode,不设置默认是200
    // w.WriteHeader(http.StatusOK)         //这行是多余的,不起作用,因为之前已经设置过响应码了
    w.Write([]byte("Hello Boy\n")) //响应体.如果Write()之前没有显式地调WriteHeader,则Write()时会先调用WriteHeader(http.StatusOK)
    // fmt.Fprint(w, "Hello Boy")
    w.Header().Add("uuid", "0987654321") //无效
    fmt.Println(strings.Repeat("*", 60))
}

func Get(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("request url: %s\n", r.URL)
    params := myhttp.ParseUrlParams(r.URL.RawQuery)
    fmt.Fprintf(w, "your anme is %s, age is %s\n", params["name"], params["age"])
    fmt.Println(strings.Repeat("*", 60))
}

func router1() {
    // 路由
    http.HandleFunc("/obs", HttpObservation)
    http.HandleFunc("/get", Get)

    // 启动Http Server
    if err := http.ListenAndServe("127.0.0.1:5678", nil); err != nil {
        panic(err)
    }
}

func main() {
    router1()
}

// go run ./http/server
