package main

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/os/gview"
    "gitee.com/johng/gf/g/net/ghttp"
    "gitee.com/johng/gf/g/util/gpage"
)

func main() {
    s := ghttp.GetServer()
    s.BindHandler("/page/demo", func(r *ghttp.Request){
        page := gpage.New(100, 10, r.Get("page"), r.URL.String())
        buffer, _ := gview.ParseContent(`
        <html>
            <head>
                <style>
                    a,span {padding:8px; font-size:16px;}
                    div{margin:5px 5px 20px 5px}
                </style>
            </head>
            <body>
                <div>{{.page1}}</div>
                <div>{{.page2}}</div>
                <div>{{.page3}}</div>
                <div>{{.page4}}</div>
            </body>
        </html>
        `, g.Map{
            "page1" : gview.HTML(page.GetContent(1)),
            "page2" : gview.HTML(page.GetContent(2)),
            "page3" : gview.HTML(page.GetContent(3)),
            "page4" : gview.HTML(page.GetContent(4)),
        })
        r.Response.Write(buffer)
    })
    s.SetPort(8199)
    s.Run()
}