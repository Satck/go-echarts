package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Kline struct {
	RectChart
}

func (Kline) Type() string {
	return types.ChartKline
}

func NewKLine() *Kline {
	c := &Kline{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

func (c *Kline) SetXAxis(xAxis interface{}) *Kline {
	c.xAxisData = xAxis
	return c
}

func (c *Kline) AddSeries(name string, data []opts.KlineData, options ...SeriesOpts) *Kline {
	series := SingleSeries{Name: name, Type: types.ChartKline, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Kline) Validata() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)

}
