package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	// 使用 r.Form 不论是 url 里的 还是 body 里的负载都可以拿到
	fmt.Fprintln(w, r.Form)

	// 使用 r.PostForm 只能拿到 请求体 body 中的负载
	// PostForm 不支持 multipart 编码的数据
	fmt.Println(r.PostForm)
}

// processMultipart
// @Desc: 	获取提交表单中 Multipart 格式编码的数据
// @Param:	w
// @Param:	r
// @Notice:
func processMultipart(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(1024)	// 表示想取1024个byte
	fmt.Println(w, r.MultipartForm)
}

// formValue
// @Desc: 	使用 r.FormValue 直接拿到 某个 key 对应的 第一个 value
// @Param:	w
// @Param:	r
// @Notice:	1. 它自动调用ParseForm或者ParseMultipartForm
//			2. 只能拿到这个键 对应的 第一个 值
func formValue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.FormValue("hello"))
}


// postFormValue
// @Desc: 	postFormValue只会返回表单中第一个值，而不管URL中的k-v
// @Param:	w
// @Param:	r
// @Notice:
func postFormValue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.PostFormValue("hello"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/mul", processMultipart)
	http.HandleFunc("/fv", formValue)
	http.HandleFunc("/pfv", postFormValue)
	server.ListenAndServe()
}
