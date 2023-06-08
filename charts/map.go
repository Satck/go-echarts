package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Map struct {
	BaseConfiguration
	BaseActions

	mapType string
}

func (Map) Type() string {
	return types.ChartMap
}

func NewMap() *Map {
	c := &Map{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

func (c *Map) AddSeries(name string, data []opts.MapData, options ...SeriesOpts) *Map {
	series := SingleSeries{Name: name, Type: types.ChartMap, MapType: c.mapType, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Map) SetGlobalOptions(options ...GlobalOpts) *Map {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

func (c *Map) SetDispatchActions(actions ...GlobalActions) *Map {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

func (c *Map) Validate() {
	c.Assets.Validate(c.AssetsHost)

}
