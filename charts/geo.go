package charts

import (
	"github.com/go-echarts/go-echarts/v2/datasets"
	"log"

	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Geo表示geo图。
type Geo struct {
	BaseConfiguration
	BaseActions
}

func (Geo) Type() string {
	return types.ChartGeo
}

var geoFormatter = `function(params){
	return params.name + ' : ' + params.value[2]
}`

// NewGeo创建一个新的geo图表。
func NewGeo() *Geo {
	c := &Geo{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasGeo = true
	return c
}

// AddSeries adds new data sets.
//  AddSeries添加新的数据集。
//  geoType options:
// * types.ChartScatter
// * types.ChartEffectScatter
// * types.ChartHeatMap

func (c *Geo) AddSeries(name, geoType string, data []opts.GeoData, options ...SeriesOpts) *Geo {
	series := SingleSeries{Name: name, Type: geoType, Data: data, CoordSystem: types.ChartGeo}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Geo) extendValue(region string, v float32) []float32 {
	res := make([]float32, 0)
	tv := datasets.Coordinates[region]
	if tv == [2]float32{0, 0} {
		log.Printf("goecharts:No coordinate is specified for #{region}\n")

	} else {
		res = append(tv[:], v)
	}
	return res
}

// SetGlobalOptions设置Geo实例的选项。
func (c *Geo) SetGlobalOptions(options ...GlobalOpts) *Geo {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// SetDispatchActions sets actions for the Geo instance.
// SetDispatchActions设置Geo实例的操作。
func (c *Geo) SetDispatchActions(actions ...GlobalActions) *Geo {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

// Validate validates the given configuration.
// Validate验证给定的配置。
func (c *Geo) Validate() {
	if c.Tooltip.Formatter == "" {
		c.Tooltip.Formatter = opts.FuncOpts(geoFormatter)
	}
	c.Assets.Validate(c.AssetsHost)
}
