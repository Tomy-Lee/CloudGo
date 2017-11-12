package server

import (
    "github.com/go-martini/martini" 
)

func NewServer(port string) {   
    m := martini.Classic()
    
    //请求处理器
    m.Get("/", func(params martini.Params) string {
        return "hey, this is Tomy, 15331151."
    })
    //创建请求端口
    m.RunOnAddr(":"+port)   
}
