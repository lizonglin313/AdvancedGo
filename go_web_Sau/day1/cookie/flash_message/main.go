package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// 使用 cookie 实现 闪现消息 flash message

// setMessage
// @Desc: 	将 消息 存入 cookie 的值
// @Param:	w
// @Param:	r
// @Notice:	对 cookie 的值进行 URL 编码 以满足浏览器要求
func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

// showMessage
// @Desc: 	实际上就是 将 上一个 同名cookie中的消息取出后
//			再将 这个 cookie 的生存时间 置为 负
// @Param:	w
// @Param:	r
// @Notice:
func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found!")
		}
	} else {
		rc := http.Cookie{
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}


}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)

	server.ListenAndServe()
}
