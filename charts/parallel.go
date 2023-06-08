package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Parallel struct {
	BaseConfiguration
	BaseActions
}

func (Parallel) Type() string {
	return types.ChartParallel
}

func NewParallel() *Parallel {
	c := &Parallel{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasParallel = true
	return c
}

func (c *Parallel) AddSeries(name string, data []opts.ParallelData, options ...SeriesOpts) *Parallel {
	series := SingleSeries{Name: name, Type: types.ChartParallel, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Parallel) SetGlobalOptions(options ...GlobalOpts) *Parallel {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

func (c *Parallel) SetDispatchActions(actions ...GlobalActions) *Parallel {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

func (c *Parallel) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
