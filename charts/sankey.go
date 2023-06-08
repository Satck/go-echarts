package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Sankey struct {
	BaseConfiguration
	BaseActions
}

func (Sankey) Type() string {
	return types.ChartSankey
}

func NewSankey() *Sankey {
	c := &Sankey{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

func (c *Sankey) AddSeries(name string, nodes []opts.SankeyNode, links []opts.SankeyLink, options ...SeriesOpts) *Sankey {
	series := SingleSeries{Name: name, Type: types.ChartSankey, Data: nodes, Links: links}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Sankey) SetGlobalOptions(options ...GlobalOpts) *Sankey {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

func (c *Sankey) SetDispatchActions(actions ...GlobalActions) *Sankey {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}
func (c *Sankey) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
