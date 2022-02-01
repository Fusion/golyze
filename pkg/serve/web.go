package serve

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	http.HandleFunc("/scripts", scriptsHelper)
	http.HandleFunc("/deps", depsPage)
	http.HandleFunc("/weight", weightPage)
	http.ListenAndServe(":8080", nil)
}

func mainPage(w http.ResponseWriter, _ *http.Request) {
	page, _ := assets.Content.ReadFile("main.tpl.html")
	fmt.Fprintf(w, "%s", page)
}

func scriptsHelper(w http.ResponseWriter, _ *http.Request) {
	content, _ := assets.Content.ReadFile("js/scripts.js")
	fmt.Fprintf(w, "%s", content)
}

func depsPage(w http.ResponseWriter, r *http.Request) {
	w_value, h_value := getWHFromRequest(r)
	fmt.Fprintf(w, "%s", renderSingleChart(
		"deps",
		charts.RenderSankey(
			ll,
			chartsData.Deps,
			w_value,
			h_value)))
}

func weightPage(w http.ResponseWriter, r *http.Request) {
	w_value, h_value := getWHFromRequest(r)
	fmt.Fprintf(w, "%s", renderSingleChart(
		"deps",
		charts.RenderPie(
			ll,
			wrap.MapToPieData(ll, chartsData.Weight),
			w_value,
			h_value)))
}

func getWHFromRequest(r *http.Request) (string, string) {
	w_idx, err := strconv.Atoi(r.URL.Query().Get("w_size"))
	if err != nil || w_idx < 0 || w_idx > 5 {
		w_idx = 0
	}
	h_idx, err := strconv.Atoi(r.URL.Query().Get("h_size"))
	if err != nil || h_idx < 0 || h_idx > 2 {
		h_idx = 0
	}
	return []string{"100%", "1024px", "1536px", "2048px", "3072px", "4096px"}[w_idx],
		[]string{"1024px", "1536px", "2048px"}[h_idx]
}
