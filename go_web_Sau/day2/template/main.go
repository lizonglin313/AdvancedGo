package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 使用 go template
// process1
// @Desc: 	判断
// @Param:	w
// @Param:	r
// @Notice:
func process1(w http.ResponseWriter, r *http.Request) {
	// ParseFiles 解析模板文件
	// 注意！！！！  /go_web_Sau/day2/template/templates/temp3.html
	// 最前面 多加一个  '/' 就出错了
	t, _ := template.ParseFiles("go_web_Sau/day2/template/templates/temp1.html")
	// 生成随机数
	rand.Seed(time.Now().Unix())
	// 将数据写到模板中
	t.Execute(w, rand.Intn(10) > 5)
}

// process2
// @Desc: 	迭代
// @Param:	w
// @Param:	r
// @Notice:
func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("go_web_Sau/day2/template/templates/temp2.html")
	numbers := []int{1, 2, 3, 4, 5}
	// numbers := []int{}   // 如果为空 触发 range 的 else
	t.Execute(w, numbers)
}

// process3
// @Desc: 	在模板里赋值
// @Param:	w
// @Param:	r
// @Notice:
func process3(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("go_web_Sau/day2/template/templates/temp3.html")
	// t := template.Must(template.ParseFiles("go_web_Sau/day2/template/templates/temp3.html"))
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "hello")
}

// process4
// @Desc: 	模板相互调用
// @Param:	w
// @Param:	r
// @Notice:
func process4(w http.ResponseWriter, r *http.Request) {
	// 第一个模板文件 被当成 主 模板
	t, err := template.ParseFiles("go_web_Sau/day2/template/templates/t1.html",
		"go_web_Sau/day2/template/templates/t2.html")
	// t := template.Must(template.ParseFiles("go_web_Sau/day2/template/templates/temp3.html"))
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "hello")
}

// process5
// @Desc: 	演示 模板变量  和 管道
// @Param:	w
// @Param:	r
// @Notice:
func process5(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("go_web_Sau/day2/template/templates/temp5.html")
	data := map[string]string{
		"Name": "lzl",
		"Age":  "22",
		"notice": "null",
	}
	t.Execute(w, data)
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

// process6
// @Desc: 	使用 FuncMap 为模板绑定函数
// @Param:	w
// @Param:	r
// @Notice:	在 模板中 可以使用 管道 、 直接在后面加 . 调用
//			使用 Key 调用
func process6(w http.ResponseWriter, r *http.Request) {
	// funcMap 绑定模板 设置 模板函数
	funcMap := template.FuncMap{"fdate" : formatDate}
	fmt.Println(time.Now())
	t := template.New("temp6.html").Funcs(funcMap)
	t,  _ = t.ParseFiles("go_web_Sau/day2/template/templates/temp6.html")
	t.Execute(w, time.Now())
}

// process7
// @Desc: 	上下文感知
// @Param:	w
// @Param:	r
// @Notice:
func process7(w http.ResponseWriter, r *http.Request) {
	t,  _ := template.ParseFiles("go_web_Sau/day2/template/templates/temp7.html")
	context := `I asked: <i>What's up?</i>'`
	t.Execute(w, context)
}

// process8
// @Desc: 	XSS 通过上下文感知 模板将 js 脚本转译成 可读字符串 而不是将他执行
// @Param:	w
// @Param:	r
// @Notice:	不能说 template 是绝对安全的 ， go template 也有方法 不转义这些 html css js的代码
func process8(w http.ResponseWriter, r *http.Request) {
	t,  _ := template.ParseFiles("go_web_Sau/day2/template/templates/temp8.html")
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("go_web_Sau/day2/template/templates/form.html")
	t.Execute(w, nil)
}

// process9
// @Desc: 	不对 html 进行转义
// @Param:	w
// @Param:	r
// @Notice:
func process9(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("X-XSS-Protection", "0")	// 有的浏览器或许自带 基本的防御XSS攻击的功能 先关掉
	t,  _ := template.ParseFiles("go_web_Sau/day2/template/templates/temp8.html")
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

// process10
// @Desc: 	使用显示定义的模板
// @Param:	w
// @Param:	r
// @Notice:
func process10(w http.ResponseWriter, r *http.Request) {
	t,  _ := template.ParseFiles("go_web_Sau/day2/template/templates/layout.html")
	// 三个参数 后两个是 使用的模板 和 数据
	t.ExecuteTemplate(w, "layout", "")
}

// process11
// @Desc: 	使用 不同文件中 的 同名 模板
// @Param:	w
// @Param:	r
// @Notice:
func process11(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template

	if rand.Intn(10)>5{
		t,  _ = template.ParseFiles("go_web_Sau/day2/template/templates/layout.html",
			"go_web_Sau/day2/template/templates/red_hello.html")
	} else {
		t,  _ = template.ParseFiles("go_web_Sau/day2/template/templates/layout.html",
			"go_web_Sau/day2/template/templates/blue_hello.html")
	}
	// 三个参数 后两个是 使用的模板 和 数据
	t.ExecuteTemplate(w, "layout", "")
}

// process12
// @Desc: 	通过定义 块动作 如果 没有达到执行 某一模板的条件 那么就执行 块动作定义的默认模板
// @Param:	w
// @Param:	r
// @Notice:
func process12(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template

	if rand.Intn(10)>5{
		t,  _ = template.ParseFiles("go_web_Sau/day2/template/templates/layout2.html",
			"go_web_Sau/day2/template/templates/red_hello.html")
	} else {
		t,  _ = template.ParseFiles("go_web_Sau/day2/template/templates/layout2.html")
	}
	// 三个参数 后两个是 使用的模板 和 数据
	t.ExecuteTemplate(w, "layout", "this data is for default template")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process1", process1)
	http.HandleFunc("/process2", process2)
	http.HandleFunc("/process3", process3)
	http.HandleFunc("/process4", process4)
	http.HandleFunc("/process5", process5)
	http.HandleFunc("/process6", process6)
	http.HandleFunc("/process7", process7)
	http.HandleFunc("/process8", process8)
	http.HandleFunc("/process9", process9)
	http.HandleFunc("/process10", process10)
	http.HandleFunc("/process11", process11)
	http.HandleFunc("/process12", process12)
	http.HandleFunc("/form", form)

	server.ListenAndServe()
}
