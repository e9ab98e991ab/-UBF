package main

import (
    "gitee.com/johng/gf/g/net/gtcp"
    "fmt"
)

func main() {
    data, err := gtcp.SendRecv("www.baidu.com:80", []byte("GET / HTTP/1.1\n\n"), -1)
    if len(data) > 0 {
        fmt.Println(string(data))
    }
    if err != nil {
        fmt.Errorf("ERROR: %s\n", err.Error())
    }
}