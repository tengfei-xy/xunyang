package doc

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"

	blackfriday "github.com/russross/blackfriday/v2"
)

// Doc of template.html
type Doc struct {
	Name        string
	Description string

	IndexShow string
	Index     template.HTML

	MDListShow string
	MDList     template.HTML

	DocsShow string
	Docs     template.HTML

	Menu template.HTML
}

// GetDocs 转化MD
func (d *Doc) GetDocs(fileLink string) {
	f, _ := ioutil.ReadFile(fileLink)
	d.Docs = template.HTML(blackfriday.Run(f))
}

// GetMenu :根据请求路径，获得当前菜单
func (d *Doc) GetMenu(currentDir string) {
	f, err := ioutil.ReadDir(currentDir)
	if err != nil {
		panic(err)
	}
	for _, i := range f {
		cate := i.Name()

		// 只允许非.开头的文件夹
		if cate[0:1] != "." && i.IsDir() {
			d.Menu += template.HTML(fmt.Sprintf(`<li class="site-menu-sub"><a href=%s>%s</li>`, cate+"/", cate))
		}
	}
}

// GetMDList :获取目录下md文件列表
func (d *Doc) GetMDList(base, fileDir string) {
	// fileDir应该是博客文件夹名称
	f, err := ioutil.ReadDir(base + fileDir)
	if err != nil {
		panic(err)
	}
	for _, i := range f {
		mdName := i.Name()
		if strings.HasSuffix(mdName, ".md") {
			d.MDList += template.HTML(fmt.Sprintf(`<div class="site-mdlist"><a href=%s>%s</div><hr class="site-mdlist-hr">`, fileDir+mdName, mdName[0:len(mdName)-3]))
		}
	}
}
