package charts

import (
	"log"

	"github.com/fusion/golyz/pkg/wrap"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func RenderSankey(l *log.Logger, data wrap.SankeyData, desiredWidth string, desiredHeight string) *charts.Sankey {
	chart := charts.NewSankey()
	chart.SetGlobalOptions(
		charts.WithInitializationOpts(
			opts.Initialization{
				AssetsHost: "https://cdn.jsdelivr.net/npm/echarts@5.3.0/dist/",
				Theme:      "white", // Could be 'dark'
				Width:      desiredWidth,
				Height:     desiredHeight,
			}),
		charts.WithTooltipOpts(
			opts.Tooltip{
				Show:      true,
				Trigger:   "item",
				TriggerOn: "mousemove",
			}),
	)
	chart.AddSeries(
		"deps",
		data.Nodes,
		data.Links,
		charts.WithEmphasisOpts(
			opts.Emphasis{
				Focus: "adjacency",
			}),
	)

	//buf := new(bytes.Buffer)
	//page := components.NewPage()
	//page.AddCharts(chart)
	//page.Render(buf)
	//chart.Render(buf)
	return chart
}
