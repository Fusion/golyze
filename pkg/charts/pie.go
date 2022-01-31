package charts

import (
	"log"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func RenderPie(l *log.Logger, pieData []opts.PieData) *charts.Pie {
	chart := charts.NewPie()
	chart.SetGlobalOptions(
		charts.WithInitializationOpts(
			opts.Initialization{
				AssetsHost: "https://cdn.jsdelivr.net/npm/echarts@5.3.0/dist/",
				Theme:      "white", // Could be 'dark'
				Width:      "1024px",
				Height:     "1024px",
			}),
		charts.WithTooltipOpts(
			opts.Tooltip{
				Show:      true,
				Trigger:   "item",
				TriggerOn: "mousemove",
			}),
	)
	chart.AddSeries(
		"weight",
		pieData,
		charts.WithEmphasisOpts(
			opts.Emphasis{
				Focus: "adjacency",
			}),
		charts.WithPieChartOpts(
			opts.PieChart{
				RoseType: "radius",
			}),
	)

	return chart
}
