package main

import (
	// golang src

	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	// my src
	"doc"
)

func xyInit() {
	log.Println("start")
}
func xyWebRun() {
	http.HandleFunc("/", xyWebIndex)
	go http.ListenAndServe("0.0.0.0:80", nil)
	log.Println(http.ListenAndServeTLS("0.0.0.0:443", "ssl/ssl.crt", "ssl/ssl.key", nil))

}
func main() {
	xyInit()
	xyWebRun()
}
func xyWebIndex(w http.ResponseWriter, r *http.Request) {
	var reqURLExt string = filepath.Ext(r.URL.Path)
	var reqURL string = r.URL.Path
	var res []byte
	var err error
	var ip string = r.RemoteAddr

	switch reqURLExt {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".ico":
		w.Header().Set("Content-Type", "image/x-ico")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".jpg":
		w.Header().Set("Content-Type", "image/jpeg")
	default:
		w.Header().Set("Content-Type", "text/html")
	}

	// 首页
	if reqURL == `/` {
		log.Printf("ip:%s - blog:%s\n", ip, reqURL)

		var d doc.Doc
		d.Name = "浔阳末栈"
		d.Description = "为学日益 为道日损"
		d.GetMenu("web/docs/mini-blog" + reqURL)
		d.IndexShow = "inline"
		d.Index = `<img src="web/static/background.jpg" width=100% height=100%">`
		d.DocsShow = "none"
		d.MDListShow = "none"
		t, _ := template.ParseFiles("web/default.html")
		t.Execute(w, d)
		return

		// 博客子目录
	} else if reqURL[len(reqURL)-1:] == `/` {
		log.Printf("ip:%s - blog directory:%s\n", ip, reqURL)

		var d doc.Doc
		d.Name = "浔阳末栈"
		d.Description = "为学日益 为道日损"
		d.GetMenu("web/docs/mini-blog" + reqURL)
		d.IndexShow = "none"
		d.DocsShow = "none"
		d.MDListShow = "inline"
		d.GetMDList("web/docs/mini-blog", reqURL)
		t, _ := template.ParseFiles("web/default.html")
		t.Execute(w, d)
		return

		// 博客文章内容
	} else if reqURLExt == ".md" {
		log.Printf("ip:%s - docs:%s\n", ip, reqURL)

		var d doc.Doc
		d.Name = "浔阳末栈"
		d.Description = "为学日益 为道日损"
		d.GetDocs("web/docs/mini-blog" + reqURL)
		d.IndexShow = "none"
		d.DocsShow = ""
		d.MDListShow = "none"
		t, _ := template.ParseFiles("web/default.html")
		t.Execute(w, d)

		// 其他文件
	} else {
		res, err = ioutil.ReadFile(reqURL[1:])
		if err != nil {
			log.Printf("ip:%s - file:%s - %v\n", ip, reqURL[1:], err)
		}
	}

	w.Write(res)

}
