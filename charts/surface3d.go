package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Surface3D struct {
	Chart3D
}

func (Surface3D) Type() string {
	return types.ChartSurface3D
}
func NewSurface3D() *Surface3D {
	c := &Surface3D{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.initChart3D()
	return c
}

func (c *Surface3D) AddSeries(name string, data []opts.Chart3DData, options ...SeriesOpts) *Surface3D {
	c.addSeries(types.ChartScatter3D, name, data, options...)
	return c
}
