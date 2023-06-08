package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Graph struct {
	BaseConfiguration
	BaseActions
}

func (Graph) Type() string {
	return types.ChartGraph
}

// NewGraph创建一个新graph表
func NewGraph() *Graph {
	chart := new(Graph)
	chart.initBaseConfiguration()
	chart.Renderer = render.NewChartRender(chart, chart.Validate)
	return chart
}

// AddSeries添加新的数据集。
func (c *Graph) AddSeries(name string, nodes []opts.GraphNode, links []opts.GraphLink, options ...SeriesOpts) *Graph {
	series := SingleSeries{Name: name, Type: types.ChartGraph, Links: links, Data: nodes}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetDispatchActions设置Graph选项的操作。
func (c *Graph) SetGlobalOptions(options ...GlobalOpts) *Graph {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// SetDispatchActions设置Graph实例的操作。
func (c *Graph) SetDispatchActions(actions ...GlobalActions) *Graph {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

// Validate验证给定的配置。
func (c *Graph) Validate() {
	for i := 0; i < len(c.MultiSeries); i++ {
		if c.MultiSeries[i].Layout == " " {
			c.MultiSeries[i].Layout = "force"
		}
	}
	c.Assets.Validate(c.AssetsHost)
}
