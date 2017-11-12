package main

import (
    "os"
	//引进我在本地文件夹下的server.go
    "web/server"
    flag "github.com/pflag"
)

const (
	//默认为8080端口
    PORT string = "8080"
)

func main() {
	//默认为8080端口
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT
    }
	//解析端口号
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
	//启动server
    server.NewServer(port)
}
