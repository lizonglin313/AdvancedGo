package main

import (
	"net/http"
)

// write01
// @Desc: 	使用 write 方法将数据写入 http Response主体中
// @Param:	w
// @Param:	r
// @Notice:
func write01(w http.ResponseWriter, r *http.Request) {

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

func re(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://lizonglin313.github.io/")
	w.WriteHeader(302)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/re", re)
	http.HandleFunc("/write", write01)
	server.ListenAndServe()
}
