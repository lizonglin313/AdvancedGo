# AdvancedGo
note&amp;code of go

go web/高级编程中的一些学习代码和笔记

## Go Web Sau

-  day0：串联处理器/处理器函数、获取请求头内容、HTTP2、http_router使用、SSL密钥证书生成

-  day1：修改响应头、cookie、文件上传、拿form数据、向Response Header/Body写数据

---

-  day2：go template的使用demo

---

-  day3：memory、file、csv、gob二进制相关的存储

---

-  day4：goroutine、channel相关内容

1. channel是可以关闭的，使用`close(c)`就可以关闭一个通道。这个时候可以使用多值来判断通道是否关闭：`c1, ok := <-a`

2. 在使用`select`选择通道时，可以用`default`处理特殊情况