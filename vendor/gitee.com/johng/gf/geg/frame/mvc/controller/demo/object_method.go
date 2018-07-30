package demo

import (
    "gitee.com/johng/gf/g/net/ghttp"
)

type ObjectMethod struct {}

func init() {
    obj := &ObjectMethod{}
    ghttp.GetServer().BindObjectMethod("/object-method", obj, "Show1, Show2, Show3")
    ghttp.GetServer().Domain("localhost").BindObjectMethod("/object-method", obj, "Show4")
}

func (o *ObjectMethod) Show1(r *ghttp.Request) {
    r.Response.Write("show 1")
}

func (o *ObjectMethod) Show2(r *ghttp.Request) {
    r.Response.Write("show 2")
}

func (o *ObjectMethod) Show3(r *ghttp.Request) {
    r.Response.Write("show 3")
}

func (o *ObjectMethod) Show4(r *ghttp.Request) {
    r.Response.Write("show 4")
}


