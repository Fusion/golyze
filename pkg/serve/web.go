package serve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fusion/golyz/pkg/assets"
	"github.com/fusion/golyz/pkg/charts"
	"github.com/fusion/golyz/pkg/data"
	"github.com/fusion/golyz/pkg/wrap"
)

var chartsData *data.ChartsData
var ll *log.Logger

func RunWebServer(l *log.Logger, dataRef *data.ChartsData) {
	l.Println("Running chart server on port 8080")

	ll = l
	chartsData = dataRef

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/deps", depsPage)
	http.HandleFunc("/weight", weightPage)
	http.ListenAndServe(":8080", nil)
}

func mainPage(w http.ResponseWriter, _ *http.Request) {
	page, _ := assets.Content.ReadFile("main.tpl.html")
	fmt.Fprintf(w, "%s", page)
}

func depsPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "%s", renderSingleChart(
		"deps",
		charts.RenderSankey(ll, chartsData.Deps)))
}

func weightPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "%s", renderSingleChart(
		"deps",
		charts.RenderPie(
			ll,
			wrap.MapToPieData(ll, chartsData.Weight))))
}
