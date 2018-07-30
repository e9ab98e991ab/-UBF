// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// CRC32
package gcrc32

import (
    "hash/crc32"
)

func EncodeString(v string) uint32 {
    return crc32.ChecksumIEEE([]byte(v))
}

func EncodeBytes(v []byte) uint32 {
    return crc32.ChecksumIEEE(v)
}
