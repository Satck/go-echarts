package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// 漏斗表示漏斗图。
type Funnel struct {
	BaseConfiguration
	BaseActions
}

// Type返回图表类型。
func (Funnel) Type() string {
	return types.ChartFunnel
}

// NewFunnel创建一个新的漏斗图。
func NewFunnel() *Funnel {
	c := &Funnel{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// SetGlobalOptions设置漏斗实例的选项。
func (c *Funnel) AddSeries(name string, data []opts.FunnelData, options ...SeriesOpts) *Funnel {
	series := SingleSeries{Name: name, Type: types.ChartFunnel, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetDispatchActions为Gauge实例设置操作。
func (c *Funnel) SetDispatchActions(actions ...GlobalActions) *Funnel {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

//Validate验证给定的配置。

func (c *Funnel) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
