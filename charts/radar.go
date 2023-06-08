package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Radar struct {
	BaseActions
	BaseConfiguration
}

func (Radar) Type() string {
	return types.ChartRadar
}

func NewRadar() *Radar {
	c := &Radar{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasRadar = true
	return c
}

func (c *Radar) AddSeries(name string, data []opts.RadarData, options ...SeriesOpts) *Radar {
	series := SingleSeries{Name: name, Type: types.ChartRadar, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	c.legends = append(c.legends, name)
	return c
}

func (c *Radar) SetGlobalOptions(options ...GlobalOpts) *Radar {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

func (c *Radar) SetDispatchActions(actions ...GlobalActions) *Radar {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

func (c *Radar) Validate() {
	c.Legend.Data = c.legends
	c.Assets.Validate(c.AssetsHost)
}
