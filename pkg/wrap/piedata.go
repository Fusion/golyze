package wrap

import (
	"fmt"
	"log"

	"github.com/dustin/go-humanize"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type PieData map[string]uint64

func MapToPieData(l *log.Logger, sourceMap map[string]uint64) []opts.PieData {
	var pieData []opts.PieData
	for moduleName, moduleSize := range sourceMap {
		data := opts.PieData{
			Name:  moduleName,
			Value: moduleSize,
			Tooltip: &opts.Tooltip{
				Show:      true,
				Formatter: fmt.Sprintf("{b}: %s", humanize.Bytes(moduleSize))},
		}
		pieData = append(pieData, data)
	}

	return pieData
}
