package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Liquid struct {
	BaseConfiguration
	BaseActions
}

func (Liquid) Type() string {
	return types.ChartLiquid
}

func NewLiquid() *Liquid {
	c := &Liquid{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.JSAssets.Add("echarts-liquidfill.min.js")
	return c
}

func (c *Liquid) AddSeries(name string, data []opts.LiquidData, options ...SeriesOpts) *Liquid {
	series := SingleSeries{Name: name, Type: types.ChartLiquid, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Liquid) SetGlobalOptions(options ...GlobalOpts) *Liquid {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

func (c *Liquid) SetDispatchActions(actions ...GlobalActions) *Liquid {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}
func (c *Liquid) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
