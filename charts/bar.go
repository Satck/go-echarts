package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// 条形图表示条形图。
type Bar struct {
	RectChart

	isXYReversal bool
}

// Type返回图表类型。
func (Bar) Type() string { return types.ChartBar }

// /NewBar创建一个新的条形图实例。
func NewBar() *Bar {
	c := &Bar{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasXYAxis = true
	return c
}

// EnablePolarType启用极坐标栏。
func (c *Bar) EnablePolarType() *Bar {
	c.hasXYAxis = false
	c.hasPolar = true
	return c
}

// SetX轴设置X轴。
func (c *Bar) SetXAxis(x interface{}) *Bar {
	c.xAxisData = x
	return c
}

// AddSeries添加新系列。
func (c *Bar) AddSeries(name string, data []opts.BarData, options ...SeriesOpts) *Bar {
	series := SingleSeries{Name: name, Type: types.ChartBar, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// XY反转检查X轴和Y轴是否反转。
func (c *Bar) XYReversal() *Bar {
	c.isXYReversal = true
	return c
}

// Validate验证给定的配置。
func (c *Bar) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	if c.isXYReversal {
		c.YAxisList[0].Data = c.xAxisData
		c.XAxisList[0].Data = nil
	}
	c.Assets.Validate(c.AssetsHost)
}
