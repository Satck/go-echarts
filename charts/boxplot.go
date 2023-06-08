package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// BoxPlot表示箱线图。
type BoxPlot struct {
	RectChart
}

// Type返回图表类型。
func (BoxPlot) Type() string {
	return types.ChartBoxPlot
}

// NewBoxPlot创建一个新的boxplot图表。
func NewBoxPlot() *BoxPlot {
	c := &BoxPlot{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

// SetX轴添加X轴。
func (c *BoxPlot) SetAxis(x interface{}) *BoxPlot {
	c.xAxisData = x
	return c
}

// AddSeries添加新系列。
func (c *BoxPlot) AddSeries(name string, data []opts.BoxPlotData, options ...SeriesOpts) *BoxPlot {
	series := SingleSeries{Name: name, Type: types.ChartBoxPlot, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate验证给定的配置。
func (c *BoxPlot) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)
}
