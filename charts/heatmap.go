package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type HeatMap struct {
	RectChart
}

func (HeatMap) Type() string {
	return types.ChartHeatMap
}

func NewHeatMap() *HeatMap {
	c := &HeatMap{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

func (c *HeatMap) SetXAxis(x interface{}) *HeatMap {
	c.xAxisData = x
	return c
}

func (c *HeatMap) AddSeries(name string, data []opts.HeatMapData, options ...SeriesOpts) *HeatMap {
	series := SingleSeries{Name: name, Type: types.ChartHeatMap, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *HeatMap) Validata() {
	c.Assets.Validate(c.AssetsHost)
}
