package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// EffectScatter表示效果散点图。
type EffectScatter struct {
	RectChart
}

// Type返回图表类型。
func (EffectScatter) Type() string {
	return types.ChartEffectScatter
}

// NewEffectScatter创建一个新的效果散点图。
func NewEffectScatter() *EffectScatter {
	c := &EffectScatter{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

// SetX轴添加X轴。
func (c *EffectScatter) SetXAxis(x interface{}) *EffectScatter {
	c.xAxisData = x
	return c
}

// AddSeries添加Y轴。
func (c *EffectScatter) AddSeries(name string, data []opts.EffectScatterData, options ...SeriesOpts) *EffectScatter {
	series := SingleSeries{Name: name, Type: types.ChartEffectScatter, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate验证给定的配置。
func (c *EffectScatter) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)
}
