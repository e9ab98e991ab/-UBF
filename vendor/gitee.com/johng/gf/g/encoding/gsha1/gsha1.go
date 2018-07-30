// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// SHA1
package gsha1

import (
    "crypto/sha1"
    "encoding/hex"
    "os"
    "io"
    "gitee.com/johng/gf/g/os/glog"
    "gitee.com/johng/gf/g/encoding/gmd5"
)

// 将任意类型的变量进行SHA摘要(注意map等非排序变量造成的不同结果)
// 内部使用了md5计算，因此效率会稍微差一些，更多情况请使用 EncodeString
func Encode(v interface{}) string {
    return EncodeString(gmd5.Encode(v))
}

// 对字符串行SHA1摘要计算
func EncodeString(s string) string {
    r := sha1.Sum([]byte(s))
    return hex.EncodeToString(r[:])
}

// 对文件内容进行SHA1摘要计算
func EncodeFile(path string) string {
    f, e := os.Open(path)
    if e != nil {
        glog.Fatalln(e)
    }
    h := sha1.New()
    _, e = io.Copy(h, f)
    if e != nil {
        glog.Fatalln(e)
    }
    return hex.EncodeToString(h.Sum(nil))
}