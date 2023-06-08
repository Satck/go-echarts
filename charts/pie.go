package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Pie struct {
	BaseConfiguration
	BaseActions
}

func (Pie) Type() string {
	return types.ChartPie
}

func NewPie() *Pie {
	c := &Pie{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

func (c *Pie) Addseries(name string, data []opts.PieData, options ...SeriesOpts) *Pie {
	series := SingleSeries{Name: name, Type: types.ChartPie, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Pie) SetGlobalOptions(options ...GlobalOpts) *Pie {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

func (c *Pie) SetDispatchActions(actions ...GlobalActions) *Pie {
	c.BaseActions.setBaseGloActions(actions...)
	return c
}

func (c *Pie) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
