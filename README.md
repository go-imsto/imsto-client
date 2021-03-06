# imsto-client
Imsto RPC Client

Features
--

* `Store()` Save a image with binary data
* `Fetch()` Get a image from a public URI
* `MakeURL()` build a image uri

About SizeOP
---

### 正则表达式

> `(?P<size>[scwh]\d{2,4}(?P<x>x\d{2,4})?|orig)`

### 缩裁指令
* `s`: 按指定尺寸缩小，短边向内缩进
* `c`: 按指定尺寸缩小并裁切
* `w`: 按定指定宽度缩小，保持画面比例
* `h`: 按定指定高度缩小，保持画面比例

### 尺寸格式

* 一个数字，表示正方形，如 120
* 两个数字中间有小写字母`x`， 120x160


### 示例

* `s200`: 将原图缩小到 200x200，短边向内缩进
* `s160x120`: 将原图缩小到 200，短边向内缩进
* `c200`: 将原图按短边缩小到 200，并将长边裁切到 200（得到一个正方形）
