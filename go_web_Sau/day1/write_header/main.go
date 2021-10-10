package main

import (
	"fmt"
	"net/http"
)

// write02
// @Desc: 	使用 write 方法将数据写入 http Response主体中
// @Param:	w
// @Param:	r
// @Notice:
func write02(w http.ResponseWriter, r *http.Request) {
	str := `
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>file_upload</title>
</head>
<body>
<form action="http://127.0.0.1:8080/processFile?hello=world&th=123"
      method="post" enctype="multipart/form-data">
    <input type="text" name="hello" value="lzl">
    <input type="text" name="post" value="456">
    <input type="file" name="uploaded">
    <input type="submit">
</form>
</body>
</html>
`
	w.Write([]byte(str))
}

// writeHeader
// @Desc: 	使用 WriteHeader 向 响应中 写入返回状态码
// @Param:	w
// @Param:	r
// @Notice:
func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintln(w, "no such service")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/write", write02)
	http.HandleFunc("/writeheader", writeHeader)
	server.ListenAndServe()
}
