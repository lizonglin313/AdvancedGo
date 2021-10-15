package main

import (
	"fmt"
	"net/http"
)

// 有关 set cookie 和 get cookie

func genCookie(n1, n2 string) (c1, c2 http.Cookie) {
	c1 = http.Cookie{
		Name:     n1,
		Value:    "Go Web",
		HttpOnly: true,
	}

	c2 = http.Cookie{
		Name:     n2,
		Value:    "Manning Go",
		HttpOnly: true,
	}
	return
}

// setCookie
// @Desc: 	第一种方法，向 Response 中写 k-v
// @Param:	w
// @Param:	r
// @Notice:
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1, c2 := genCookie("first_cookie", "second_cookie")
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}


// setCookie2
// @Desc: 	第二种方法 使用 http 库中的 SetCookie 方法直接设置
// @Param:	w
// @Param:	r
// @Notice:
func setCookie2(w http.ResponseWriter, r *http.Request) {
	c1, c2 := genCookie("3th_cookie", "4th_cookie")
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

// getCookie
// @Desc: 	方法一 从 request header 中拿到
// @Param:	w
// @Param:	r
// @Notice:
func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header["Cookie"]
	fmt.Fprintln(w, cookie)
}

// getCookie2
// @Desc: 	第二种方法 使用 request 的 Cookie 和 Cookies 方法，从请求中拿
// @Param:	w
// @Param:	r
// @Notice:	Cookie("cookie_name")只能拿到1个 Cookies()可以拿到所有的
func getCookie2(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")	// r.Cookie 返回一个 cookie 指针
	if err != nil {
		fmt.Fprintln(w, "Cannot get first_cookie")
	}

	cs := r.Cookies()	// r.Cookies 返回一个 cookie 的指针切片
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/set_cookie2", setCookie2)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/get_cookie2", getCookie2)

	server.ListenAndServe()
}
