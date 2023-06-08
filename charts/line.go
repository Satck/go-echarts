package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Line struct {
	RectChart
}

func (Line) Type() string {
	return types.ChartLine
}

func NewLine() *Line {
	c := &Line{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}
func (c *Line) SetXAxis(x interface{}) *Line {
	c.xAxisData = x
	return c
}

func (c *Line) AddSeries(name string, data []opts.LineData, options ...SeriesOpts) *Line {
	series := SingleSeries{Name: name, Type: types.ChartLine, Data: data}
	series.InitSeriesDefaultOpts(c.BaseConfiguration)
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Line) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)

}
