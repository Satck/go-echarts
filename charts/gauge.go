package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Gauge 表示仪表聊天
type Gauge struct {
	BaseConfiguration
	BaseActions
}

// Type返回图表类型。
func (Gauge) Type() string {
	return types.ChartGauge
}

// NewGauge 创建新的gague图表
func NewGauge() *Gauge {
	c := &Gauge{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// AddSeries添加新的数据集
func (c *Gauge) AddSeries(name string, data []opts.GaugeData, options ...SeriesOpts) *Gauge {
	series := SingleSeries{Name: name, Type: types.ChartGauge, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions设置Gauge实例的选项
func (c *Gauge) SetGlobalOptions(options ...GlobalOpts) *Gauge {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// SetDispatchActions设置Gauge安装的操作
func (c *Gauge) SetDispatchActions(actions ...GlobalActions) *Gauge {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

// Validate验证给定的配置
func (c *Gauge) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
