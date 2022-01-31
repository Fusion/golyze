package serve

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"

	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/render"

	"github.com/fusion/golyz/pkg/assets"
)

var (
	pat = regexp.MustCompile(`(__f__")|("__f__)|(__f__)`)
)

func renderSingleChart(name string, singleChart components.Charter) []byte {
	page := &components.Page{}
	page.Assets.InitAssets()
	page.Layout = components.PageCenterLayout

	page.AddCharts(singleChart)

	inittpl, _ := assets.Content.ReadFile("init.tpl.html")
	basetpl, _ := assets.Content.ReadFile("base.tpl.html")
	htmltpl, _ := assets.Content.ReadFile("singlechart.tpl.html")
	tpl := template.Must(template.New("deps").Parse(string(inittpl))).Funcs(template.FuncMap{
		"safeJS": func(s interface{}) template.JS {
			return template.JS(fmt.Sprint(s))
		},
	})
	for _, t := range [][]byte{basetpl, htmltpl} {
		tpl = template.Must(tpl.Parse(string(t)))
	}

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, render.ModPage, page); err != nil {
		ll.Fatal(err)
	}

	return pat.ReplaceAll(buf.Bytes(), []byte(""))
}

/*
type flexPageRender struct {
	c      interface{}
	before []func()
}

func NewFlexRenderer() *components.Page {
	page := &components.Page{}
	page.Assets.InitAssets()
	page.Renderer = NewPageRender(page, page.Validate)
	page.Layout = components.PageCenterLayout
	return page
}

func NewPageRender(c interface{}, before ...func()) render.Renderer {
	return &flexPageRender{c: c, before: before}
}

func (r *flexPageRender) Render(w io.Writer) error {
	for _, fn := range r.before {
		fn()
	}

	headertpl, _ := assets.Content.ReadFile("header.tpl.html")
	basetpl, _ := assets.Content.ReadFile("base.tpl.html")
	pagetpl, _ := assets.Content.ReadFile("page.tpl.html")
	contents := []string{
		string(headertpl),
		string(basetpl),
		string(pagetpl)}
	tpl := render.MustTemplate(render.ModPage, contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, render.ModPage, r.c); err != nil {
		return err
	}

	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := w.Write(content)
	return err
}
*/
