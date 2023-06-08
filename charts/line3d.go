package charts

import "C"
import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Line3D struct {
	Chart3D
}

func (Line3D) Type() string {
	return types.ChartLine3D
}

func NewLine3D() *Line3D {
	c := &Line3D{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.initChart3D()
	return c
}

func (c *Line3D) AddSeries(name string, data []opts.Chart3DData, options ...SeriesOpts) *Line3D {
	c.addSeries(types.ChartLine3D, name, data, options...)
	return c
}
