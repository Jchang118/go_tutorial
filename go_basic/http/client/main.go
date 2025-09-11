package main

import (
    //"bufio"
    //"bytes"
    //"encoding/json"
    "fmt"
    "io"
    "net/http"
    //"net/url"
    "os"
    //"strconv"
    "strings"
    //"time"

    myhttp "go_tutorial/go_basic/http"
)

// 详细看一下http协议
func HttpObservation() {
    fmt.Println(strings.Repeat("*", 30) + "GET" + strings.Repeat("*", 30))
    if resp, err := http.Get("http://127.0.0.1:5678/obs?name=dqq"); err != nil {
        panic(err)
    } else {
        defer resp.Body.Close() //注意:一定要调用resp.Body.Close(),否则会协程泄漏(同时引发内存泄漏)
        fmt.Printf("response proto: %s\n", resp.Proto)
        if major, minor, ok := http.ParseHTTPVersion(resp.Proto); ok {
            fmt.Printf("http major version %d, http minor version %d\n", major, minor)
        }

        fmt.Printf("response status: %s\n", resp.Status)
        fmt.Println("response header")
        for key, values := range resp.Header {
            fmt.Printf("%s: %v\n", key, values)
            if key == "Date" {
                if tm, err := http.ParseTime(values[0]); err == nil {
                    fmt.Printf("server time %s\n", tm.Format("2006-01-02 15:04:05"))
                }
            }
        }
        fmt.Println("response body:")
        io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
        os.Stdout.WriteString("\n\n")
    }
}

func Get() {
    fmt.Println(strings.Repeat("*", 30) + "GET" + strings.Repeat("*", 30))
    if resp, err := http.Get("http://127.0.0.1:5678/get?" + myhttp.EncodeUrlParams(map[string]string{"name": "良诚 Chang", "age": "18"})); err != nil {
        panic(err)
    } else {
        defer resp.Body.Close()
        fmt.Printf("response status: %s\n", resp.Status)
        fmt.Println("response body:")
        // io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
        if body, err := io.ReadAll(resp.Body); err == nil {
            fmt.Print(string(body))
        }
        os.Stdout.WriteString("\n\n")
    }
}

func main() {
    HttpObservation()
    Get()
}

// go run ./http/client
