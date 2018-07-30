// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// MD5
package gmd5

import (
    "crypto/md5"
    "fmt"
    "encoding/json"
    "reflect"
    "os"
    "io"
    "gitee.com/johng/gf/g/os/glog"
)

// 将任意类型的变量进行md5摘要(注意map等非排序变量造成的不同结果)
func Encode(v interface{}) string {
    h := md5.New()
    if "string" == reflect.TypeOf(v).String() {
        h.Write([]byte(v.(string)))
    } else {
        b, err := json.Marshal(v)
        if err != nil {
            return ""
        } else {
            h.Write(b)
        }
    }
    return fmt.Sprintf("%x", h.Sum(nil))
}

// 将字符串进行MD5哈希摘要计算
func EncodeString(v string) string {
    h := md5.New()
    h.Write([]byte(v))
    return fmt.Sprintf("%x", h.Sum(nil))
}

// 将文件内容进行MD5哈希摘要计算
func EncodeFile(path string) string {
    f, e := os.Open(path)
    if e != nil {
        glog.Fatalln(e)
    }
    h := md5.New()
    _, e = io.Copy(h, f)
    if e != nil {
        glog.Fatalln(e)
    }
    return fmt.Sprintf("%x", h.Sum(nil))
}
