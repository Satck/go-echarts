package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type ThemeRiver struct {
	BaseActions
	BaseConfiguration
}

func (ThemeRiver) Type() string {
	return types.ChartThemeRiver
}
func NewThemeRiver() *ThemeRiver {
	c := &ThemeRiver{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasSingleAxis = true
	return c
}

func (c *ThemeRiver) AddSeries(name string, data []opts.ThemeRiverData, options ...SeriesOpts) *ThemeRiver {
	cd := make([][3]interface{}, len(data))
	for i := 0; i < len(data); i++ {
		cd[i] = data[i].ToList()
	}
	series := SingleSeries{Name: name, Type: types.ChartThemeRiver, Data: cd}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *ThemeRiver) SetGlobalOptions(options ...GlobalOpts) *ThemeRiver {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}
func (c *ThemeRiver) SetDispatchActions(actions ...GlobalActions) *ThemeRiver {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}
func (c *ThemeRiver) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
